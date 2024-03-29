import 'package:comies/components.dart';
import 'package:comies/components/shared/default-button.dart';
import 'package:comies/core.dart';
import 'package:comies/products/products.dart';
import 'package:datacontext/datacontext.dart';
import 'package:flutter/material.dart';

class IngredientForm extends StatefulWidget {
  final Ingredient? ingredient;

  final LoadStatus submitStatus;

  final Function() onSubmit;

  final bool isNew;

  const IngredientForm({Key? key, required this.ingredient, required this.onSubmit, required this.submitStatus, this.isNew = true}) : super(key: key);

  @override
  _IngredientFormState createState() => _IngredientFormState();
}

class _IngredientFormState extends State<IngredientForm> {
  late TextEditingController quantityController;
  late TextEditingController ingredientController;
  late TextEditingController nameController;
  late GlobalKey<FormState> formState;

  void setFields() {
    quantityController = TextEditingController(text: getDoubleView(widget.ingredient?.quantity, 2))..addListener(() => widget.ingredient?.quantity = getDoubleValue(quantityController.text));
    ingredientController = TextEditingController(text: getTextView(widget.ingredient?.ingredientId))..addListener(() => widget.ingredient?.ingredientId = getTextValue(ingredientController.text));
    nameController = TextEditingController(text: getTextView(widget.ingredient?.component?.name));
  }

  @override
  void initState() {
    super.initState();
    formState = GlobalKey<FormState>();
    setFields();
  }

  @override
  void didUpdateWidget(IngredientForm oldWidget) {
    super.didUpdateWidget(oldWidget);
    setFields();
  }

  @override
  Widget build(BuildContext context) {
    return Form(
      key: formState,
      child: Padding(
        padding: const EdgeInsets.all(8.0),
        child: Column(
          children: [
            if (widget.ingredient!.component != null)
              Padding(
                padding: const EdgeInsets.only(bottom: 10.0),
                child: Text('Produto: ' + nameController.text, style: getMainText(size: 22)),
              ),
            Row(
              children: [
                Center(
                  child: DefaultIconButton(
                    label: 'Diminuir',
                    icon: Icons.remove,
                    onTap: () => quantityController.text = getDoubleView(getDoubleValue(quantityController.text) - 1, 2),
                  ),
                ),
                Expanded(
                  child: TextFormField(
                    controller: quantityController,
                    decoration: InputDecoration(labelText: 'Quantidade'),
                  ),
                ),
                Center(
                  child: DefaultIconButton(
                    label: 'Aumentar',
                    icon: Icons.add,
                    onTap: () => quantityController.text = getDoubleView(getDoubleValue(quantityController.text) + 1, 2),
                  ),
                ),
              ],
            ),
            Expanded(
              child: Container(
                height: getHeight(context) * 0.8,
                child: LoadStatusWidget(
                  status: getProvider(context).productViews.loadStatus,
                  loadWidget: (context) => ProductSelector(
                    data: getProvider(context).productViews.local['selectorResults'] ?? <ProductView>[],
                    onSearch: (f) => ProductsController.searchProductsToLocalStorage(context, f, dataIdentifier: 'selectorResults'),
                    onCancelSearch: () => getProvider(context).productViews.local['selectorResults'] = <ProductView>[],
                    onTap: (p) {
                      ingredientController.text = p.id;
                      setState(() {
                        nameController.text = p.name;
                      });
                    },
                  ),
                ),
              ),
            ),
            MainButton(
              label: 'Salvar ingrediente',
              loadingLabel: 'Salvando ingrediente...',
              onTap: () {
                if (formState.currentState!.validate()) {
                  widget.onSubmit();
                }
              },
            ),
          ],
        ),
      ),
    );
  }
}
