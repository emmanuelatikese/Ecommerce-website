
import EmptyCart from "../components/emptyCart"
import OrderSummary from "../components/order";
import CartCards from './../components/cartCards';


const Cart = () => {
  const IsEmpty = false
  return (
    <div className="flex flex-col items-center ">
      <h1 className="text-4xl font-bold font-custom pt-4">Cart</h1>
      {IsEmpty ? <EmptyCart/> :

        <div className="flex flex-row  w-11/12">
            <div className=" flex flex-col justify-center items-center  w-10/12 ">
              <CartCards/>
            </div>
            <OrderSummary/>
        </div>
      }
    </div>
  )
}

export default Cart