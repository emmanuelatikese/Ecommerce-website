import Slider from "../components/slider"

const Home = () => {

  return (
    <div className="flex flex-col justify-center items-center  pt-4 gap-4">
    <div className="flex flex-col gap-1">
        <h1 className="text-4xl font-bold font-custom pt-4">EXPLORE OUR CATEGORIES</h1>
        <p className="text-dash text-center font-sans">Discover the latest trend in an eco-friendly fashion</p>
    </div>
      <Slider/>

      <div className="flex flex-col gap-1">
        <h1 className="text-4xl font-bold font-custom pt-4">FEATURED PRODUCTS</h1>
    </div>


    </div>
  )
}

export default Home
