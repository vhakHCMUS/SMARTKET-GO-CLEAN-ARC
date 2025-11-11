package handlers

import (
	"net/http"
	"strconv"
	"github.com/vhakHCMUS/SMARTKET-GO-CLEAN-ARC/domains/order"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	orderService order.Service
}

// NewOrderHandler creates a new order handler
func NewOrderHandler(orderService order.Service) *OrderHandler {
	return &OrderHandler{orderService: orderService}
}

// CreateOrder creates a new order
// @Summary Create an order
// @Tags orders
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param request body order.CreateOrderRequest true "Order details"
// @Success 201 {object} order.Order
// @Router /api/orders [post]
func (h *OrderHandler) CreateOrder(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
		return
	}

	var req order.CreateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ord, err := h.orderService.CreateOrder(userID.(uint), &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": ord})
}

// GetOrder gets an order by ID
// @Summary Get order details
// @Tags orders
// @Security BearerAuth
// @Produce json
// @Param id path int true "Order ID"
// @Success 200 {object} order.Order
// @Router /api/orders/{id} [get]
func (h *OrderHandler) GetOrder(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid order id"})
		return
	}

	ord, err := h.orderService.GetOrderByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "order not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": ord})
}

// GetUserOrders gets all orders for the authenticated user
// @Summary Get user's orders
// @Tags orders
// @Security BearerAuth
// @Produce json
// @Success 200 {array} order.Order
// @Router /api/orders [get]
func (h *OrderHandler) GetUserOrders(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
		return
	}

	orders, err := h.orderService.GetUserOrders(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": orders})
}

// GetMerchantOrders gets all orders for the merchant
// @Summary Get merchant's orders
// @Tags orders
// @Security BearerAuth
// @Produce json
// @Success 200 {array} order.Order
// @Router /api/merchant/orders [get]
func (h *OrderHandler) GetMerchantOrders(c *gin.Context) {
	merchantID, exists := c.Get("merchantID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "merchant not authenticated"})
		return
	}

	orders, err := h.orderService.GetMerchantOrders(merchantID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": orders})
}

// RedeemOrder redeems an order (merchant confirms pickup)
// @Summary Redeem/confirm order
// @Tags orders
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param request body order.RedeemOrderRequest true "Order code"
// @Success 200
// @Router /api/merchant/orders/redeem [post]
func (h *OrderHandler) RedeemOrder(c *gin.Context) {
	merchantID, exists := c.Get("merchantID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "merchant not authenticated"})
		return
	}

	var req order.RedeemOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.orderService.RedeemOrder(merchantID.(uint), req.OrderCode); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "order redeemed successfully"})
}

// AddToCart adds an item to cart
// @Summary Add item to cart
// @Tags cart
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param request body order.AddToCartRequest true "Cart item"
// @Success 200
// @Router /api/cart/add [post]
func (h *OrderHandler) AddToCart(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
		return
	}

	var req order.AddToCartRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.orderService.AddToCart(userID.(uint), &req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "item added to cart"})
}

// GetCart gets user's cart
// @Summary Get cart
// @Tags cart
// @Security BearerAuth
// @Produce json
// @Success 200 {object} order.Cart
// @Router /api/cart [get]
func (h *OrderHandler) GetCart(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
		return
	}

	cart, err := h.orderService.GetCart(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": cart})
}

// UpdateCartItem updates cart item quantity
// @Summary Update cart item
// @Tags cart
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "Cart Item ID"
// @Param request body order.UpdateCartItemRequest true "Quantity"
// @Success 200
// @Router /api/cart/items/{id} [put]
func (h *OrderHandler) UpdateCartItem(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid item id"})
		return
	}

	var req order.UpdateCartItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.orderService.UpdateCartItem(userID.(uint), uint(id), req.Quantity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "cart item updated"})
}

// RemoveCartItem removes an item from cart
// @Summary Remove cart item
// @Tags cart
// @Security BearerAuth
// @Param id path int true "Cart Item ID"
// @Success 200
// @Router /api/cart/items/{id} [delete]
func (h *OrderHandler) RemoveCartItem(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid item id"})
		return
	}

	if err := h.orderService.RemoveCartItem(userID.(uint), uint(id)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "cart item removed"})
}

// ClearCart clears user's cart
// @Summary Clear cart
// @Tags cart
// @Security BearerAuth
// @Success 200
// @Router /api/cart/clear [post]
func (h *OrderHandler) ClearCart(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
		return
	}

	if err := h.orderService.ClearCart(userID.(uint)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "cart cleared"})
}
