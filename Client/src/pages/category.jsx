
import CatCards from './../components/categoryCard';

const Category = () => {
  return (
      <div className='flex flex-col items-center gap-10'>  
        <h1 className="text-4xl font-bold font-custom pt-2">Clothes Category</h1> 
            <div className='flex flex-row items-center flex-wrap gap-4 justify-center'>
            <CatCards/>
            <CatCards/>
            <CatCards/>
            <CatCards/>
            <CatCards/>
            <CatCards/>
            </div>
      </div>
  )
}

export default Category