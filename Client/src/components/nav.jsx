import { FaHome } from "react-icons/fa";
import { FaShoppingCart } from "react-icons/fa";
import { Link } from "react-router-dom";
import { IoLogOutSharp } from "react-icons/io5";
import { IoLogIn   } from "react-icons/io5";

const nav = () => {
    const User = true
  return (
    <div className="p-4 w-full flex px-7 border-b-2 border-x-0 border-t-0 flex-wrap shadow-sm shadow-marigold-200 border-b-marigold-200 border-4 flex-row justify-between items-center text-marigold-200">
    <p className="font-bold flex"><p className="text-white">Ur.</p>UsualMarket</p>

    <div className="flex flex-row gap-4">
    {
        User && (
            <Link>
                    <FaHome className="w-6 h-6"/>
            </Link>
        )
    }

    {
        User && (
        <Link>
        <FaShoppingCart className="w-6 h-6"/>
        </Link>
        )
    }


 { User?  <p className="flex justify-center text-white text-sm w-24 font-bold border-solid border-2 p-1 border-marigold-200 rounded-md"><Link>Dashboard</Link></p> :
    <p className="flex justify-center text-white text-sm w-24 font-bold border-solid border-2 p-1 border-marigold-200 rounded-md"><Link to={"/SignUp"}>SignUp</Link></p> }

{User ? <Link>
                <IoLogOutSharp className="w-6 h-6"/>
        </Link> :
    <Link to={"/Login"}>
            <IoLogIn className="w-6 h-6"/>
    </Link>}
        
    </div>

    </div>
 )
}

export default nav
