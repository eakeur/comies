import { useEffect } from 'react';
import { Home } from './components/home';
import { useShortcuts } from './hooks/keypress';
import { addAction } from './handlers/handler';
import { useDrawer } from './hooks/drawer';
import { ProductForm } from './components/product/product-form';
import { Text } from '@chakra-ui/react'
import { createBrowserRouter, RouterProvider } from 'react-router-dom';
import { PlaceOrder } from 'components/orders/place';

const router = createBrowserRouter([
  {
    path: "/",
    element: <Home/>,
  },
]);


function App() {
  const shortcuts = useShortcuts()
  const drawer = useDrawer()

  useEffect(() => {

    addAction({
      name:"add_product", 
      handler: () => {
        drawer({
          header: <Text as="h1">Novo produto</Text>,
          body: <ProductForm/>
        })
      }
    })

    addAction({
      name:"add_order", 
      handler: () => {
        drawer({
          header: <Text as="h1">Novo pedido</Text>,
          body: <PlaceOrder/>,
          options: {
            size: "xl"
          }
        })
      }
    })
  }, [shortcuts, drawer])

  return <RouterProvider router={router}/>
}

export default App;
