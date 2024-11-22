import CategoryCards from "../components/categoryCards"

const Category = () => {
  return (
      <div className='flex flex-col items-center'>  
        <h1 className="text-4xl font-bold font-custom pt-2">Clothes Category</h1> 
            <div className='flex flex-row items-center flex-wrap gap-4'>
            <CategoryCards/>
            </div>
      </div>
  )
}

export default Category