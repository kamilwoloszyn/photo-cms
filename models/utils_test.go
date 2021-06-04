package models_test

import (
	"time"

	"github.com/google/uuid"
	"github.com/kamilwoloszyn/photo-cms/configs"
	"github.com/kamilwoloszyn/photo-cms/models"
	"github.com/kamilwoloszyn/photo-cms/pkg/database"
	. "github.com/onsi/ginkgo"
	"github.com/pkg/errors"
)

var _ = Describe("Util testing", func() {

})

func ConnectToDatabase() error {
	dbConfig, err := configs.LoadDbConfig()
	if err != nil {
		return err
	}
	db, err := database.Initialize(*dbConfig)
	if err != nil {
		return err
	}
	if err := models.SetHandler(db); err != nil {
		return err
	}
	return nil
}

func CreateUserId() uuid.UUID {
	return uuid.New()
}

func CreateOptionValue(o *models.Option) models.OptionValue {
	sampleValue1 := models.OptionValue{
		Value:      "Yellow",
		ExtraPrice: 10.0,
		OptionId:   o.GetID(),
	}
	if err := sampleValue1.Create(); err != nil {
		errWrapped := errors.Wrap(err, "Creating Option value1")
		Fail(errWrapped.Error())

	}
	return sampleValue1
}

func CreateCustomOptionValue(o *models.Option, ov *models.OptionValue) {
	ov.OptionId = o.GetID()
	if err := ov.Create(); err != nil {
		errWrapped := errors.Wrap(err, "Creating custom option value")
		Fail(errWrapped.Error())
	}
}

func CreateOption() models.Option {
	sampleOption := models.Option{
		Name: "Color",
	}
	if err := sampleOption.Create(); err != nil {
		wrappedErr := errors.Wrap(err, "Creating Option")
		Fail(wrappedErr.Error())
	}
	return sampleOption
}

func CreateCustomOption(opt *models.Option) {
	if err := opt.Create(); err != nil {
		errWrapped := errors.Wrap(err, "Creating custom option")
		Fail(errWrapped.Error())
	}
}

func CreateCategory() models.Category {
	sampleCategory := models.Category{
		CategoryName: "Odbitki",
	}
	if err := sampleCategory.Create(); err != nil {
		errWrapped := errors.Wrap(err, "Creating sample category")
		Fail(errWrapped.Error())
	}
	return sampleCategory
}

func CreateCustomCategory(c *models.Category) {
	if err := c.Create(); err != nil {
		errWrapped := errors.Wrap(err, "Creating custom category")
		Fail(errWrapped.Error())
	}
}

func CreateProductOption(p *models.Product, v *models.OptionValue) models.ProductOption {
	productOption := models.ProductOption{
		ProductId:     p.GetID(),
		OptionValueId: v.GetID(),
	}
	if err := productOption.Create(); err != nil {
		errWrapped := errors.Wrap(err, "Creating product option")
		Fail(errWrapped.Error())
	}
	return productOption
}

func CreateCustomProductOption(p *models.Product, ov *models.OptionValue, po *models.ProductOption) {
	po.OptionValueId = ov.GetID()
	po.ProductId = p.GetID()

	if err := po.Create(); err != nil {
		errWrapped := errors.Wrap(err, "Creating custom product option")
		Fail(errWrapped.Error())
	}
}
func CreateImage() models.Image {
	sampleImage := models.Image{
		Name:     "generated_img",
		FullPath: "/tmp/generated_img",
		Size:     10000,
	}
	if err := sampleImage.Create(); err != nil {
		errWrapped := errors.Wrap(err, "Creating new image")
		Fail(errWrapped.Error())
	}
	return sampleImage
}

func CreateCustomImage(i *models.Image) {
	if err := i.Create(); err != nil {
		errWrapped := errors.Wrap(err, "Creating custom image")
		Fail(errWrapped.Error())
	}
}

func CreateDeliveryMethod() models.DeliveryMethod {
	deliveryMethod := models.DeliveryMethod{
		Name:       "InPost",
		FixedPirce: 8.99,
	}
	if err := deliveryMethod.Create(); err != nil {
		errWrapped := errors.Wrap(err, "Creating delivery Methods")
		Fail(errWrapped.Error())
	}
	return deliveryMethod
}

func CreateCustomDeliveryMethod(d *models.DeliveryMethod) {
	if err := d.Create(); err != nil {
		errWrapped := errors.Wrap(err, "Creating custom delivery method")
		Fail(errWrapped.Error())
	}
}

func CreatePaymentMethod() models.PaymentMethod {
	paymentMethod := models.PaymentMethod{
		Name:        "PayU",
		Provider:    "PayU",
		PosId:       "57139243",
		KeyMd5:      "15117b282328146ac6afebaa8acd80e7",
		ClientId:    "768246287",
		OauthSecret: "15117b282328146a6affecea8acdw0e7",
	}
	if err := paymentMethod.Create(); err != nil {
		errWrapped := errors.Wrap(err, "Create Payment Method")
		Fail(errWrapped.Error())
	}
	return paymentMethod
}

func CreateCustomPaymentMethod(p *models.PaymentMethod) {
	if err := p.Create(); err != nil {
		errWrapped := errors.Wrap(err, "Creating custom payment method")
		Fail(errWrapped.Error())
	}
}

func CreatePayment(pm *models.PaymentMethod) models.Payment {
	t := time.Now()
	payment := models.Payment{
		PaymentDate:     &t,
		PaymentAmount:   300.11,
		PaymentMethodId: pm.GetID(),
		PaymentError:    false,
		PaymentFinished: false,
	}
	if err := payment.Create(); err != nil {
		errWrapped := errors.Wrap(err, "Crating payment")
		Fail(errWrapped.Error())
	}
	return payment
}
func CreateCustomPayment(pm *models.PaymentMethod, p *models.Payment) {
	time := time.Now()
	p.PaymentMethodId = pm.GetID()
	p.PaymentDate = &time
	if err := p.Create(); err != nil {
		errWrapped := errors.Wrap(err, "Creating custom payment")
		Fail(errWrapped.Error())
	}
}

func CreateDelivery(dm *models.DeliveryMethod) models.Delivery {
	delivery := models.Delivery{
		ShippedVia:               "Michal",
		TrackingCode:             "123793472742342",
		DestinationPostalCode:    "37-630",
		DestinationConturyRegion: "Podkarpackie",
		DestinationAddress:       "Zamkowa 100/10",
		DestinationCity:          "Oleszyce",
		DeliveryMethodId:         dm.GetID(),
	}
	if err := delivery.Create(); err != nil {
		errWrapped := errors.Wrap(err, "Creating delivery")
		Fail(errWrapped.Error())
	}
	return delivery
}
func CreateCustomDelivery(dm *models.DeliveryMethod, d *models.Delivery) {
	d.DeliveryMethodId = dm.GetID()
	if err := d.Create(); err != nil {
		errWrapped := errors.Wrap(err, "Creating custom delivery")
		Fail(errWrapped.Error())
	}
}

func CreateCustomer() models.Customer {
	customer := models.Customer{
		City:         "Oleszyce",
		Address:      "Zamkowa 100/10",
		LastName:     "Kowalski",
		FirstName:    "Michal",
		PostalCode:   "37-630",
		CompanyName:  "ABB",
		PhoneNumber:  "123456789",
		EmailAddress: "name@exmaple.com",
		Employed:     false,
		NIP:          "",
		Regon:        "",
	}
	if err := customer.Create(); err != nil {
		errWrapped := errors.Wrap(err, "Creating customer")
		Fail(errWrapped.Error())
	}
	return customer
}

func CreateCustomCustomer(c *models.Customer) {
	if err := c.Create(); err != nil {
		errWrapped := errors.Wrap(err, "Creating custom customer")
		Fail(errWrapped.Error())
	}
}

func CreateProductWithoutOrder(category *models.Category, image *models.Image, customer *models.Customer) models.Product {
	product := models.Product{
		UnitPrice:   0,
		ProductName: "sample_image",
		CategoryId:  category.GetID(),
		ImageId:     image.GetID(),
		CustomerId:  customer.GetID(),
		Quantity:    3,
	}
	if err := product.Create(); err != nil {
		errWrapped := errors.Wrap(err, "Creating product without order")
		Fail(errWrapped.Error())
	}
	return product
}
func CreateCustomProductWithoutOrder(category *models.Category, image *models.Image, customer *models.Customer, p *models.Product) {
	p.CategoryId = category.GetID()
	p.ImageId = image.GetID()
	p.CustomerId = customer.GetID()

	if err := p.Create(); err != nil {
		errWrapped := errors.Wrap(err, "Creating custom category withou order")
		Fail(errWrapped.Error())
	}
}

func CreateOrder(p *models.Payment, d *models.Delivery) models.Order {
	order := models.Order{
		Fvat:       true,
		Price:      320,
		PaymentId:  p.GetID(),
		DeliveryId: d.GetID(),
	}
	if err := order.Create(); err != nil {
		errWarpped := errors.Wrap(err, "Creating order")
		Fail(errWarpped.Error())
	}
	return order
}

func CreateCustomOrder(p *models.Payment, d *models.Delivery, o *models.Order) {
	o.PaymentId = p.GetID()
	o.DeliveryId = d.GetID()

	if err := o.Create(); err != nil {
		errWrapped := errors.Wrap(err, "Creating custom order")
		Fail(errWrapped.Error())
	}
}

func CreateProductWithOrder(c *models.Category, i *models.Image, customer *models.Customer, o *models.Order) models.Product {
	product := models.Product{
		UnitPrice:   0,
		ProductName: "sample_image",
		CategoryId:  c.GetID(),
		ImageId:     i.GetID(),
		CustomerId:  customer.GetID(),
		OrderId:     o.ID,
		Quantity:    3,
	}
	if err := product.Create(); err != nil {
		errWrapped := errors.Wrap(err, "Creating product with order")
		Fail(errWrapped.Error())
	}
	return product
}

func CreateCustomProductWithOrder(c *models.Category, i *models.Image, cs *models.Customer, o *models.Order, p *models.Product) {
	p.CategoryId = c.GetID()
	p.ImageId = i.GetID()
	p.CustomerId = cs.GetID()
	p.OrderId = o.GetID()

	if err := p.Create(); err != nil {
		errWrapped := errors.Wrap(err, "Creating custom product without order")
		Fail(errWrapped.Error())
	}
}
