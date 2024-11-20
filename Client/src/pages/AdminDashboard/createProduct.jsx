import AdminBtn from "../../components/adminBtns";
import CreateProdCards from "../../components/createProdCards";



const CreateProduct = () => {
  return (
    <div className="flex flex-col items-center gap-3">
    <h1 className="text-4xl font-bold font-custom pt-2">ADMIN DASHBOARD</h1>
    <AdminBtn/>

    <CreateProdCards/>

    </div>
  )
}

export default CreateProduct
