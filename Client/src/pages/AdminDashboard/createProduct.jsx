import AdminBtn from "../../components/adminBtns";
import { LuUpload } from "react-icons/lu";


const CreateProduct = () => {
  return (
    <div className="flex flex-col items-center gap-8">
    <h1 className="text-4xl font-bold font-custom pt-8">ADMIN DASHBOARD</h1>
    <AdminBtn/>

    <div className="absolute top-72 p-2 w-96 flex flex-col gap-4 bg-white rounded shadow-md items-center">
    <h1 className="font-bold font-custom text-color3">CREATE PRODUCT </h1>
    <div className="w-80 flex flex-col gap-4">

      <div className="flex flex-col gap-1 justify-start">
      <p className="font-custom text-dash text-color3 ">Product Name</p>
      <input type="text"  className="bg-color4 text-dash p-2 pl-6 outline-none rounded-lg shadow-md font-bold"/>
      </div>


      <div className="flex flex-col gap-1 justify-start">
      <p className="font-custom text-dash text-color3">Description</p>
      <textarea type="text"  className="bg-color4 w-80 text-dash p-2 pl-6 outline-none rounded-lg shadow-md font-bold"/>
      </div>

      <div className="flex flex-col gap-1 justify-start">
      <p className="font-custom text-dash text-color3">Price</p>
      <input type="number"  className="text-dash bg-color4 p-2 pl-6 outline-none rounded-lg shadow-md font-bold"/>
      </div>

      <div className="flex flex-col gap-1 justify-start">
      <p className="font-custom text-dash text-color3">Category</p>
      <select className="text-dash bg-color4 p-2 pl-6 outline-none rounded-lg shadow-md font-bold">
      <option value="" disabled selected className="outline-none" >Select a category</option>
      <option className="outline-none"></option>
      </select>
      </div>

      <div className="flex flex-col gap-1 justify-start">
      <p className="font-custom text-dash text-color3">Quantity</p>
      <input type="number"  className="text-dash bg-color4 p-2 pl-6 outline-none rounded-lg shadow-md font-bold"/>
      </div>

      <div className="flex flex-col gap-1 justify-start w-28">
      <p className="font-custom text-dash text-color3">Image upload</p>
      <input type="file" id="image" className="sr-only" accept="image/*" />

      <label htmlFor="image" className="cursor-pointer w-40 font-bold flex flex-row items-center gap-2  text-white bg-color3 rounded-md shadow-md p-2 text-dash">
      <LuUpload className="w-8 h-8"/>
       <p>Upload Image</p>
      </label>
      </div>
      <button className="font-custom text-white bg-color2 p-2 rounded-sm"> create</button>
    </div>
    </div>

    </div>
  )
}

export default CreateProduct
