package booking

import "github.com/gin-gonic/gin"

type BookingController struct {
	bookingService *BookingService
}

func NewBookingController(router *gin.Engine, bookingService *BookingService) BookingController {
	controller := BookingController{
		bookingService: bookingService,
	}
	router.Handle("POST", "/api/v1/bookings", controller.CreateBooking)
	router.Handle("GET", "/api/v1/bookings/:id", controller.GetBookingById)
	router.Handle("DELETE", "/api/v1/bookings/:id", controller.CancelBooking)

	return controller
}

func (c *BookingController) CreateBooking(ctx *gin.Context) {

}

func (c *BookingController) GetBookingById(ctx *gin.Context) {}

func (c *BookingController) CancelBooking(ctx *gin.Context) {}
