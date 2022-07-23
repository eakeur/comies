package app

import (
	"comies/app/core/entities/ingredient"
	"comies/app/core/entities/item"
	"comies/app/core/entities/movement"
	"comies/app/core/entities/order"
	"comies/app/core/entities/product"
	"comies/app/core/workflows/menu"
	"comies/app/core/workflows/ordering"
	"comies/app/gateway/persistence/postgres/transaction"
	"comies/app/sdk/id"
	"go.uber.org/zap"

	"github.com/bwmarrin/snowflake"
	"github.com/jackc/pgx/v4/pgxpool"
)

type (
	Gateways struct {
		Logger        *zap.SugaredLogger
		Database      *pgxpool.Pool
		SnowflakeNode *snowflake.Node
	}

	Actions struct {
		Products    product.Actions
		Ingredients ingredient.Actions
		Movements   movement.Actions
		Orders      order.Actions
		Items       item.Actions
	}

	Managers struct {
		Logger       *zap.SugaredLogger
		Transactions transaction.Manager
		ID           id.Manager
	}

	Application struct {
		Managers Managers
		Menu     menu.Workflow
		Ordering ordering.Workflow
	}
)
