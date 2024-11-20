import AdminBtn from "../../components/adminBtns"
import Trouser from "../../assets/images/trouser.jpg"

const AllProduct = () => {
  return (
    <div className="flex flex-col items-center gap-3">
    <h1 className="text-4xl font-bold font-custom pt-2">ADMIN DASHBOARD</h1>
    <AdminBtn/>

    <div className="flex flex-col gap-4 items-center">
    
    <div className="flex flex-row items-start ">
    <img src={Trouser} className="w-60 h-60"/>
    <div className="flex flex-col gap-2">
    
    </div>
    </div>
    
    </div>


    </div>
  )
}

export default AllProduct