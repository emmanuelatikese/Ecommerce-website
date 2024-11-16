
import Trouser from "../assets/images/trouser.jpg"
import { ImPriceTags } from "react-icons/im";
import { HiMiniRectangleStack } from "react-icons/hi2";
import { PiShoppingCartSimpleFill } from "react-icons/pi";

const FeaturedCard = () => {
  return (
          <div className="w-62 shadow-md shadow-slate-500 rounded-lg h-auto p-2 flex flex-col gap-0 items-center">
            <img src={Trouser} className=" w-56"/>
            <div className="w-56 bg-white p-5 flex flex-col gap-4">
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
                                <p>VIEW </p>
                                </div>
              </div>

          </div>
  )
}

export default FeaturedCard
