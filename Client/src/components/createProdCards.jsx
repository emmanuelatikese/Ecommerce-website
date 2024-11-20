import { LuUpload } from "react-icons/lu";
import { HiMiniRectangleStack } from 'react-icons/hi2';
import { IoMdPricetag } from "react-icons/io";
import { BiSolidDetail } from "react-icons/bi";
import { ImSortAmountDesc } from "react-icons/im";

const CreateProdCards = () => {
  return (
<div className="absolute top-56 p-8 w-96 flex flex-col gap-2 bg-white rounded shadow-md items-center">
    <h1 className="font-bold font-custom text-color3">CREATE PRODUCT </h1>
    <div className="w-80 flex flex-col gap-6">

      <div className="flex flex-row gap-2 items-center pl-4 bg-color4 rounded-md shadow-md">
       <HiMiniRectangleStack className="w-6 text-color3"/>
      <input type="text" placeholder="Product name" className="w-72 bg-color4 text-dash p-2 outline-none font-bold"/>
      </div>

      <div className="flex flex-row gap-2 items-center pl-4  bg-color4 rounded-md shadow-md">
       <BiSolidDetail className="w-6 text-color3"/>
            <textarea type="text" placeholder="Description" className="w-72 bg-color4 text-dash p-2 outline-none font-bold"/>
      </div>

      <div className="flex flex-row gap-2 items-center pl-4  bg-color4 rounded-md shadow-md">
       <IoMdPricetag className="w-6 text-color3"/>
      <input type="number" step={0.01} placeholder="Price" className=" bg-color4 text-dash p-2 w-72 outline-none font-bold"/>
      </div>

      <div className="flex flex-col gap-1 justify-start">
      <select  className="text-dash bg-color4 p-2 pl-6 outline-none rounded-lg shadow-md font-bold">
      <option value="" disabled selected className="outline-none" > Select a category</option>
      <option className="outline-none"></option>
      </select>
      </div>

      <div className="flex flex-row gap-2 items-center pl-4  bg-color4 rounded-md shadow-md">
       <ImSortAmountDesc className="w-6 text-color3"/>
      <input type="number"  placeholder="Quantity" className=" bg-color4 text-dash p-2 w-72 outline-none font-bold"/>
      </div>

      <div className="flex flex-col gap-1 justify-start w-28">
      <input type="file" id="image" className="sr-only" accept="image/*" />

      <label htmlFor="image" className="p-2 cursor-pointer w-36 font-bold flex flex-row items-center gap-4  text-white bg-color3 rounded-md shadow-md  text-dash">
      <LuUpload className="w-4 h-4"/>
       <p>Upload Image</p>
      </label>
      </div>
      <button className="font-custom text-white bg-color2 p-2 rounded-sm"> create</button>
    </div>
    </div>
  )
}

export default CreateProdCards
