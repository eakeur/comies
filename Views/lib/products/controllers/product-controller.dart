import 'package:comies/core.dart';
import 'package:datacontext/datacontext.dart';
import 'package:flutter/material.dart';

class ProductController {
  final BuildContext context;

  final String id;

  late DataSet<Product> products;

  late DataSet<Ingredient> ingredients;

  late DataSet<Stock> stock;

  late DataSet<StockMovement> movements;

  ProductController(this.context, [this.id = 'new']) {
    try {
      products = getProvider(context).products.replicate();
      ingredients = products.rel<Ingredient>('ingredients', parentId: id);
      stock = getProvider(context).stocks.replicate();
      movements = stock.rel<StockMovement>('stockMovements', parentId: id);
    } catch (e) {
      print(e);
    }
  }

  void get() {
    if (id == 'new') {
      products.create(Product(creationDate: DateTime.now()));
    } else {
      products.getOne(id).then((value) => ingredients.get().then((value) => stock.getOne(id).then((value) => movements.get()))).catchError((e) {
        showErrorFeedback(e, context);
      });
    }
  }

  void add() => launchFuture(products.add(), context, false, () => Navigator.pushReplacementNamed(context, 'products/${products.data!.id}', arguments: products.data!.name));
  void update() => launchFuture(products.update(id), context);
  void delete() => launchFuture(products.remove(id), context);

  void addIngredient() => launchFuture(ingredients.add(ingredients.local['editingIngredient']), context, true);
  void updateIngredient() => launchFuture(ingredients.update(ingredients.data!.id), context, true);
  void deleteIngredient([String? id]) => launchFuture(ingredients.remove(id ?? ingredients.data!.id), context);

  void addMovements() => launchFuture(movements.add(movements.local['editingMovements']), context, true);
  void updateMovements() => launchFuture(movements.update(movements.data!.id), context, true);
  void deleteMovements([String? id]) => launchFuture(movements.remove(id ?? movements.data!.id), context);
}
