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

func CreateUserId() string {
	return uuid.NewString()
}

func CreateOptionValues() []models.OptionValue {
	sampleValue1 := models.OptionValue{
		Value:      "Black",
		ExtraPrice: 10.0,
	}

	sampleValue2 := models.OptionValue{
		Value:      "Red",
		ExtraPrice: 11.0,
	}

	sampleValue3 := models.OptionValue{
		Value:      "White",
		ExtraPrice: 12.0,
	}
	if err := sampleValue1.Create(); err != nil {
		errWrapped := errors.Wrap(err, "Creating Option value1")
		Fail(errWrapped.Error())
	}

	if err := sampleValue2.Create(); err != nil {
		errWrapped := errors.Wrap(err, "Creating Option value2")
		Fail(errWrapped.Error())
	}

	if err := sampleValue3.Create(); err != nil {
		errWrapped := errors.Wrap(err, "Creating Option value3")
		Fail(errWrapped.Error())
	}
	return []models.OptionValue{sampleValue1, sampleValue2, sampleValue3}
}

func CreateOptionValue() models.OptionValue {
	sampleValue1 := models.OptionValue{
		Value:      "Yellow",
		ExtraPrice: 10.0,
	}
	if err := sampleValue1.Create(); err != nil {
		errWrapped := errors.Wrap(err, "Creating Option value1")
		Fail(errWrapped.Error())

	}
	return sampleValue1
}

func CreateOption(opt *[]models.OptionValue) models.Option {
	sampleOption := models.Option{
		Name:          "Color",
		OptionsValues: *opt,
	}
	if err := sampleOption.Create(); err != nil {
		wrappedErr := errors.Wrap(err, "Creating Option")
		Fail(wrappedErr.Error())
	}
	return sampleOption
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
func CreateProductOption(p *models.Product, v *[]models.OptionValue) models.ProductOption {
	productOption := models.ProductOption{
		Products: []models.Product{
			*p,
		},
		OptionValues: *v,
	}
	if err := productOption.Create(); err != nil {
		errWrapped := errors.Wrap(err, "Creating product option")
		Fail(errWrapped.Error())
	}
	return productOption
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

func CreatePayment(pm *[]models.PaymentMethod) models.Payment {
	t := time.Now()
	payment := models.Payment{
		PaymentDate:     &t,
		PaymentAmount:   300.11,
		PaymentMethods:  *pm,
		PaymentError:    false,
		PaymentFinished: false,
	}
	if err := payment.Create(); err != nil {
		errWrapped := errors.Wrap(err, "Crating payment")
		Fail(errWrapped.Error())
	}
	return payment
}

func CreateDelivery(dm *[]models.DeliveryMethod) models.Delivery {
	delivery := models.Delivery{
		ShippedVia:               "Michal",
		TrackingCode:             "123793472742342",
		DestinationPostalCode:    "37-630",
		DestinationConturyRegion: "Podkarpackie",
		DestinationAddress:       "Zamkowa 100/10",
		DestinationCity:          "Oleszyce",
		DeliveryMethods:          *dm,
	}
	if err := delivery.Create(); err != nil {
		errWrapped := errors.Wrap(err, "Creating delivery")
		Fail(errWrapped.Error())
	}
	return delivery
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
		NIP:          "123-456-779",
		Regon:        "1234324",
	}
	if err := customer.Create(); err != nil {
		errWrapped := errors.Wrap(err, "Creating customer")
		Fail(errWrapped.Error())
	}
	return customer
}

func CreateEmployedCustomer() models.Customer {
	customer := models.Customer{
		City:         "Oleszyce",
		Address:      "Zamkowa 100/11",
		LastName:     "Kowalski",
		FirstName:    "Jan",
		PostalCode:   "37-630",
		CompanyName:  "ABB",
		PhoneNumber:  "123456799",
		EmailAddress: "master@exmaple.com",
		Employed:     true,
		NIP:          "",
		Regon:        "",
	}
	if err := customer.Create(); err != nil {
		errWrapped := errors.Wrap(err, "Creating employed customer")
		Fail(errWrapped.Error())
	}
	return customer
}
func CreateProductWithoutOrder(c *[]models.Category, i *[]models.Image, customer *[]models.Customer) models.Product {
	product := models.Product{
		UnitPrice:   0,
		ProductName: "sample_image",
		Category:    *c,
		Image:       *i,
		Customer:    *customer,
		Quantity:    3,
	}
	if err := product.Create(); err != nil {
		errWrapped := errors.Wrap(err, "Creating product without order")
		Fail(errWrapped.Error())
	}
	return product
}

func CreateOrder(p *models.Payment, d *models.Delivery) models.Order {
	order := models.Order{
		Fvat:     true,
		Price:    320,
		Payment:  *p,
		Delivery: *d,
	}
	if err := order.Create(); err != nil {
		errWarpped := errors.Wrap(err, "Creating order")
		Fail(errWarpped.Error())
	}
	return order
}

func CreateProductWithOrder(c *[]models.Category, i *[]models.Image, customer *[]models.Customer, o *models.Order) models.Product {
	product := models.Product{
		UnitPrice:   0,
		ProductName: "sample_image",
		Category:    *c,
		Image:       *i,
		Customer:    *customer,
		Quantity:    3,
		Order:       *o,
	}
	if err := product.Create(); err != nil {
		errWrapped := errors.Wrap(err, "Creating product with order")
		Fail(errWrapped.Error())
	}
	return product
}
