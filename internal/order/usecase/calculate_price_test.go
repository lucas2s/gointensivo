package usecase

import (
	"database/sql"
	"testing"

	"github.com/lucas2s/gointensivo/internal/order/Infra/database"
	"github.com/lucas2s/gointensivo/internal/order/entity"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

type CalculateFinalPriceTestSuit struct {
	suite.Suite
	OrderRepository database.OrderRepository
	Db              *sql.DB
}

func (suite *CalculateFinalPriceTestSuit) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	suite.NoError(err)
	_, err = db.Exec("CREATE TABLE orders (id VARCHAR(255) NOT NULL, price float NOT NULL, tax float NOT NULL, final_price float NOT NULL, PRIMARY KEY (id))")
	suite.NoError(err)
	suite.Db = db
	suite.OrderRepository = *database.NewOrderRepository(db)
}

func (suite *CalculateFinalPriceTestSuit) TearDownTest() {
	suite.Db.Close()
}

func TestSuit(t *testing.T) {
	suite.Run(t, new(CalculateFinalPriceTestSuit))
}

func (suite *CalculateFinalPriceTestSuit) TestCalculateFinalPrice() {
	order, err := entity.NewOrder("123", 10.0, 2.0)
	suite.NoError(err)
	suite.NoError(order.CalculateFinalPrice())
	calculateFinalPriceInput := OrderInputDTO{
		ID:    order.ID,
		Price: order.Price,
		Tax:   order.Tax,
	}

	CalculateFinalPriceUseCase := NewCalculateFinalPriceUseCase(suite.OrderRepository)
	calculateFinalPriceOutput, err := CalculateFinalPriceUseCase.Execute(calculateFinalPriceInput)
	suite.NoError(err)
	suite.Equal(order.ID, calculateFinalPriceOutput.ID)
	suite.Equal(order.Price, calculateFinalPriceOutput.Price)
	suite.Equal(order.Tax, calculateFinalPriceOutput.Tax)
	suite.Equal(order.FinalPrice, calculateFinalPriceOutput.FinalPrice)
}
