import { HiMiniRectangleStack } from "react-icons/hi2"
import Trouser from "../assets/images/trouser.jpg"
import { PiShoppingCartSimpleFill } from "react-icons/pi"
import { ImPriceTags } from "react-icons/im"



const CatCards = () => {
  return (
    <div className=" shadow-md flex flex-col p-4 rounded-md gap-2 items-center">
        <img  src={Trouser} className="w-64 rounded-md shadow-md"/>
        <div className=" bg-white p-5 flex flex-col gap-2 rounded-md shadow-md w-40 h-36">
                  <div className="flex flex-row items-center gap-5">
                              <HiMiniRectangleStack className="w-4 h-4 text-color3"/>
                              <p className="font-custom ">TROUSER</p>
                  </div>

                  <div className="flex flex-row items-center gap-5">
                              <ImPriceTags className="w-4 h-4 text-color2"/>
                              <p className="font-custom text-color2">$400</p>
                  </div>
                                <div className="cursor-pointer bg-color2 rounded-sm shadow-md p-2 flex flex-row items-center text-white font-bold gap-4 w-auto h-10"> 
                                <PiShoppingCartSimpleFill className="w-8 h-8 text-white"/>
                                <p className="text-sm">ADD TO CART</p>
                                </div>
              </div>
    </div>
  )
}

export default CatCards