import AdminBtn from "../../components/adminBtns"
import AllProductCards from "../../components/allProductCards"


const AllProduct = () => {
  return (
    <div className="flex flex-col items-center gap-3">
    <h1 className="text-4xl font-bold font-custom pt-2">ADMIN DASHBOARD</h1>
    <AdminBtn/>

    <div className="flex  flex-row flex-wrap gap-8 items-start p-4">
    
    <AllProductCards />
    <AllProductCards />
    <AllProductCards />
    <AllProductCards />
    
    
    </div>


    </div>
  )
}

export default AllProduct