import { Link } from "react-router-dom";
import EmptyCart from "../components/emptyCart"
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
        <CartCards/>
        <CartCards/>
        <CartCards/>
        <CartCards/>
        <CartCards/>
        <CartCards/>
      </div>

      <div className="fixed right-10  flex flex-col gap-4 flex-wrap">
        
          <div className="bg-white flex flex-col w-80 p-8 gap-4 rounded-md shadow-md">
            <h3 className="font-custom">Order Summary</h3>
          
            <div className="flex flex-row justify-between">
              <p>Original price</p>
              <p>$180</p>
            </div>
          <hr className="text-color4"/>

          <div className="flex flex-row justify-between">
              <h3 className="font-custom">Total price</h3>
              <p className="font-custom text-green-600">$180</p>
            </div>

            <button className="p-2 bg-color2 text-white font-bold rounded-md shadow-sm">Proceed to Checkout</button>
            <Link className="underline font-bold flex justify-center text-green-700">
            <p>Continue to shopping</p>
            
            </Link> 
          
            </div>


      
<div className="bg-white flex flex-col w-80 p-8 gap-4  rounded-md shadow-md">
            <h3 className="text-dash font-bold "> Do you have a voucher or a gift card?</h3>
          
            <input type="text" placeholder="Enter your coupon code here...." className="bg-color4 p-2 text-dash rounded-md shadow-md outline-none font-bold"/>

            <button className="p-2 bg-color2 text-white font-bold rounded-md shadow-sm">Apply Code</button>
            <div className="flex flex-col gap-1">
            <p className="text-dash font-bold text-gray-600">Your available coupon:</p>
              <p className="text-dash font-bold text-gray-600">XHGJFHE235443 - 10% off </p>
            </div>
              
            </div>

      </div>
        </div>

      }
    </div>
  )
}

export default Cart