import AdminBtn from "../../components/adminBtns"
import Trouser from "../../assets/images/trouser.jpg"
import { CiStar } from "react-icons/ci";
import { MdDelete } from "react-icons/md";

const AllProduct = () => {
  return (
    <div className="flex flex-col items-center gap-3">
    <h1 className="text-4xl font-bold font-custom pt-2">ADMIN DASHBOARD</h1>
    <AdminBtn/>

    <div className="flex flex-col gap-4 items-center p-4 bg-white rounded-md shadow-md">
    
    <div className="flex flex-row items-start w-96 gap-4">
    <img src={Trouser} className="w-60 h-60"/>
    <div className="flex flex-col gap-2 p-2 items-center">
        <p className="font-custom text-dash">Trouser</p>
        <p className="font-serif text-color2 font-extrabold">$400</p>
        <p className="font-bold text-color3"> Cloth</p>
        <div className="flex flex-row gap-3">
          < CiStar className="w-6 h-6 text-color3"/>
          <MdDelete className="w-6 h-6 text-red-600" />
        </div>

    </div>
    </div>
    
    </div>


    </div>
  )
}

export default AllProduct