import { PiShoppingCartSimpleFill } from "react-icons/pi";
import { Link } from 'react-router-dom';


const EmptyCart = () => {
  return (
      <div className="flex flex-col justify-center items-center pt-40 gap-2">
        <PiShoppingCartSimpleFill className="w-32 h-32"/>
        <p className="font-bold text-2xl">Your cart is empty</p>
        <p className="text-dash text-gray-700 ">It seems you have{"n't"} added any item to your cart yet</p>
        <Link>
        <button className="bg-color2 p-4 text-white rounded-md shadow-md font-bold">
        Start shopping
        </button>
        </Link>
      </div>
  )
}

export default EmptyCart
