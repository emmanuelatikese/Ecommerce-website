import { HiMiniRectangleStack } from "react-icons/hi2"
import Trouser from "../assets/images/trouser.jpg"
import { PiShoppingCartSimpleFill } from "react-icons/pi"
import { ImPriceTags } from "react-icons/im"



const ProdCards = () => {
  return (
    <div className="bg-color4 shadow-md flex flex-col p-4 rounded-md gap-4">
        <img  src={Trouser} className="w-60 rounded-md shadow-md"/>
        <div className=" bg-white p-5 flex flex-col gap-4 rounded-md shadow-md">
                  <div className="flex flex-row items-center gap-5">
                              <HiMiniRectangleStack className="w-8 h-8 text-color3"/>
                              <p className="font-custom ">TROUSER</p>
                  </div>

                  <div className="flex flex-row items-center gap-5">
                              <ImPriceTags className="w-8 h-8 text-color2"/>
                              <p className="font-custom text-color2">$400</p>
                  </div>
                                <div className="cursor-pointer border-color2 border-2 p-2 flex flex-row items-center text-color2 font-bold gap-8"> 
                                <PiShoppingCartSimpleFill className="w-8 h-8 text-color2"/>
                                <p className="text-dash">ADD TO CART</p>
                                </div>
              </div>
    </div>
  )
}

export default ProdCards