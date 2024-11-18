import AdminBtn from "../../components/adminBtns"


const CreateProduct = () => {
  return (
    <div className="flex flex-col items-center gap-8">
    <h1 className="text-4xl font-bold font-custom pt-8">ADMIN DASHBOARD</h1>
    <AdminBtn/>

    <div className="absolute top-72 p-2 w-96 flex flex-col gap-4 bg-white rounded shadow-md items-center">
    <h1 className="font-bold font-custom text-color3">CREATE PRODUCT </h1>
    
    </div>

    </div>
  )
}

export default CreateProduct
