import EmptyCart from "../components/emptyCart"
import Trouser from "../assets/images/trouser.jpg"
import { MdDelete } from "react-icons/md";

const Cart = () => {
  const IsEmpty = false
  return (
    <div className="flex flex-col items-center ">
      <h1 className="text-4xl font-bold font-custom pt-4">Cart</h1>
      {IsEmpty && <EmptyCart/>}

      <div className="pt-12 flex flex-col justify-center items-center  w-10/12"> 
      
      <div className=" p-4 round-md shadow-md flex flex-row gap-12">
      <div className="bg-color4 p-4 shadow-md rounded-md"><img src={Trouser} className="w-36 h-36"/></div>
      <div className="bg-color4 p-4 flex flex-row gap-4 rounded-md shadow-md">
        <div className="flex flex-col gap-2 p-2 justify-center ">
        <p className="font-custom ">Trousers</p>
        <p className="font-bold text-gray-600">Cloth</p>
        <MdDelete className="w-6 h-6 text-red-600"/>
        </div>

        <div>
          
        </div>
      </div>
      </div>
      
      </div>
      
    </div>
  )
}

export default Cart