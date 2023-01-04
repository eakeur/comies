import { randFood, randNumber } from "@ngneat/falso"
import { rest } from "msw"


const { faker } = require("http-request-mock")

export const handlers = [
  rest.get("/api/v1/ordering/orders/:order_id", (req, res, ctx) => {
    const phone = req.url.searchParams.get("phone")

    if (phone === null) {
      return res(ctx.status(404))
    }
    return res(
      ctx.json({
        value: faker.rand(1, 5),
        occurred_at: new Date(new Date().getTime() - 60000 * faker.rand(0, 60)),
        order_id: "8787984849",
        customer_name: faker.name(),
      }),
      ctx.delay(faker.rand(0, 2000)),
      ctx.status(200),
    )
  }),

  rest.get("/api/v1/ordering/orders", (req, res, ctx) => {
    const status = req.url.searchParams.get("statuscount")

    if (status === "") {
      return res(ctx.json([]))
    }

    return res(
      ctx.delay(faker.rand(0, 2000)),
      ctx.json({count: faker.rand(0, 15)}),
      ctx.status(200),
    )
  }),

  rest.get("api/v1/menu/items/saleable", (req, res, ctx) => {
    const code = req.url.searchParams.get("code") ?? ""
    const len = () : number => 20 - code.length * 2 <= 0 ? 1 : 20 - code.length * 2

    return res(
      ctx.json(randFood({ length: len() }).map(f => ({
        id: faker.rand(1000, 100000),
        name: f,
        code: f.toUpperCase().split(" ").at(0)?.substring(0, f.split(" ").at(0)?.length ?? 0 <= 6 ? undefined : 6),
        price: randNumber({ min: 500, max: 15000}),
        stock: randNumber({ min: 0, max: 1000}),
      })))      
    )
  })
]