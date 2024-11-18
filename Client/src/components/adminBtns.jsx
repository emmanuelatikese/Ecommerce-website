import { Link } from "react-router-dom"
import { IoMdAdd } from "react-icons/io";
import { SiGoogleanalytics } from "react-icons/si";
import { HiMiniRectangleStack } from "react-icons/hi2";

const AdminBtn = () => {
  return (
    <div className="flex flex-row gap-4 items-center absolute top-48">
      
        <Link>
        <button className=" bg-color2 flex items-center gap-4 p-4 text-white rounded-md shadow-md">
        <IoMdAdd className="w-8 h-8"/>
        <p className="font-bold">CREATE PRODUCT</p>
        </button>
        </Link>

        <Link>
        <button className=" bg-white flex items-center gap-4 p-4 text-black rounded-md shadow-md">
        <HiMiniRectangleStack className="w-8 h-8"/>
        <p className="font-bold">ALL PRODUCTS</p>
        </button>
        </Link>

        <Link>
        <button className=" bg-color3 flex items-center gap-4 p-4 text-white rounded-md shadow-md">
        <SiGoogleanalytics className="w-8 h-8"/>
        <p className="font-bold">ANALYTICS</p>
        </button>
        </Link>

    </div>
  )
}

export default AdminBtn
