import Slider from "../components/slider";
const home = () => {
  return (
    <div className="text-marigold-200 flex flex-col justify-center items-center ">
      <h1 className=" font-bold text-4xl pt-12">Explore Our Categories</h1>
      <p className="text-white font-bold pt-2">Discover the latest trends</p>
      <Slider/>
      </div>
  )
}

export default home