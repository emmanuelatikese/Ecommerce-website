import Trouser from "../assets/images/trouser.jpg"
import { CiStar } from "react-icons/ci";
import { MdDelete } from "react-icons/md";
import { ImPriceTags } from "react-icons/im";
import { BiSolidCategory } from "react-icons/bi";

const AllProductCards = () => {
  return (
<div className="flex flex-row items-start ">
    <img src={Trouser} className="w-60 h-60 rounded-md shadow-md rounded-tr-none"/>
    <div className="flex flex-col gap-2 p-4 rounded-l-none rounded-md shadow-md items-center w-40 bg-white">
        <p className="font-custom text-dash">Trouser</p>
        <div className="flex flex-row gap-4">
            <div className="flex flex-row items-center gap-1">
          <ImPriceTags className="w-4 h-4 text-color2"/>
        <p className="font-san text-color2 font-extrabold">
        $400</p>
    </div>
    <div className="flex flex-row items-center gap-1">
          <BiSolidCategory className="w-4 h-4 text-color3"/>
        <p className="font-san text-dash text-color3 font-extrabold">
        Cloth</p>
    </div>
        </div>


        <div className="flex flex-row gap-3">
          <CiStar className="w-6 h-6 text-color3"/>
          <MdDelete className="w-6 h-6 text-red-600" />
        </div>

    </div>
    </div>
  )
}

export default AllProductCards