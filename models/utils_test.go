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

func CreateSomeOptionValues(o *models.Option) []models.OptionValue {
	sampleValue1 := models.OptionValue{
		Value:      "Red",
		ExtraPrice: 10.0,
		OptionId:   o.GetID(),
	}
	sampleValue2 := models.OptionValue{
		Value:      "Green",
		ExtraPrice: 10.0,
		OptionId:   o.GetID(),
	}
	sampleValue3 := models.OptionValue{
		Value:      "Blue",
		ExtraPrice: 10.0,
		OptionId:   o.GetID(),
	}

	if err := sampleValue1.Create(); err != nil {
		errWrapped := errors.Wrap(err, "Creating option value1")
		Fail(errWrapped.Error())
	}

	if err := sampleValue2.Create(); err != nil {
		errWrapped := errors.Wrap(err, "Creating option value2")
		Fail(errWrapped.Error())
	}

	if err := sampleValue3.Create(); err != nil {
		errWrapped := errors.Wrap(err, "Creating option value3")
		Fail(errWrapped.Error())
	}

	return []models.OptionValue{
		sampleValue1,
		sampleValue2,
		sampleValue3,
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

func CreateSomeDeliveryMethods() []models.DeliveryMethod {
	deliveryMethod1 := models.DeliveryMethod{
		Name:       "InPost",
		FixedPirce: 8.99,
	}
	deliveryMethod2 := models.DeliveryMethod{
		Name:       "InPost",
		FixedPirce: 8.99,
	}

	deliveryMethod3 := models.DeliveryMethod{
		Name:       "InPost",
		FixedPirce: 8.99,
	}

	if err := deliveryMethod1.Create(); err != nil {
		Fail(err.Error())
	}
	if err := deliveryMethod2.Create(); err != nil {
		Fail(err.Error())
	}
	if err := deliveryMethod3.Create(); err != nil {
		Fail(err.Error())
	}
	return []models.DeliveryMethod{
		deliveryMethod1,
		deliveryMethod2,
		deliveryMethod3,
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

func CreateSomePaymentMethods() []models.PaymentMethod {
	paymentMethod1 := models.PaymentMethod{
		Name:        "PayUp",
		Provider:    "PayUp",
		PosId:       "57169243",
		KeyMd5:      "15117b282328146ac6afebaa8acd80e7",
		ClientId:    "768246287",
		OauthSecret: "15111b282328646a6affecea8acdw0e7",
	}
	paymentMethod2 := models.PaymentMethod{
		Name:        "PayPal",
		Provider:    "PayPal",
		PosId:       "572313404",
		KeyMd5:      "151162342328146ac6afebaa8acd80e7",
		ClientId:    "768245587",
		OauthSecret: "15117b282328146a6affeaea8acdw0e7",
	}
	paymentMethod3 := models.PaymentMethod{
		Name:        "Przelewy24",
		Provider:    "Przelewy24",
		PosId:       "57139260",
		KeyMd5:      "15167b282328346ac6afebaa8acd80e7",
		ClientId:    "768246687",
		OauthSecret: "15116b282028146a6affecea8acdw0e7",
	}
	if err := paymentMethod1.Create(); err != nil {
		errWrapped := errors.Wrap(err, "Creating paymentmethod1")
		Fail(errWrapped.Error())
	}
	if err := paymentMethod2.Create(); err != nil {
		errWrapped := errors.Wrap(err, "Creating paymentmethod2")
		Fail(errWrapped.Error())
	}
	if err := paymentMethod3.Create(); err != nil {
		errWrapped := errors.Wrap(err, "Creating epaymentmethod3")
		Fail(errWrapped.Error())
	}

	return []models.PaymentMethod{
		paymentMethod1,
		paymentMethod2,
		paymentMethod3,
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
		NIP:          "123-456-779",
		Regon:        "1234324",
	}
	if err := customer.Create(); err != nil {
		errWrapped := errors.Wrap(err, "Creating employed customer")
		Fail(errWrapped.Error())
	}
	return customer
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
