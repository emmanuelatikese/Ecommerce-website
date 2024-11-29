import EmptyCart from "../components/emptyCart"

const Cart = () => {
  const IsEmpty = false
  return (
    <div className="flex flex-col items-center ">
      <h1 className="text-4xl font-bold font-custom pt-4">Cart</h1>
      {IsEmpty && <EmptyCart/>}

      <div className="flex flex-col justify-center items-center "> 
      
      
      
      </div>
      
    </div>
  )
}

export default Cart