import { Link } from "react-router-dom";
import Home from './../pages/home';
import { MdHomeFilled } from "react-icons/md";
import { PiShoppingCartSimpleFill } from "react-icons/pi";

const Nav = () => {
  return (
    <div
      className="w-full sticky bg-white h-auto rounded-md p-4 shadow-slate-400 shadow-md flex flex-row flex-wrap justify-between items-center"
    >
    <h3 className="font-sans">E-commerce</h3>

    <div className="flex flex-row gap-8">

    <Link to={Home}>
    <MdHomeFilled className="w-8 h-8 bg-white text-color3 shadow-md shadow-color1 rounded"/>
    </Link>

    <Link to={Home} className="flex flex-row flex-wrap w-14 h-8 pl-1 bg-white shadow-md shadow-color1 rounded">
    <PiShoppingCartSimpleFill className="w-8 h-8 text-color3 "/>
    <div className="w-4 h-4 mt-3 mr-0.5 text-center text-white  text-sm  bg-color2 rounded-2xl"> 5</div>
    </Link>

    <div className="bg-color2  text-dash  font-bold text-center p-2 text-white rounded-md shadow"> Dashboard </div>
    </div>
    </div>
  )
}

export default Nav