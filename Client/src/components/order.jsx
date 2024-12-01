import { Link } from "react-router-dom";
const OrderSummary = () => {
  return (
    <div>
      <div className="fixed right-10  flex flex-col gap-4 flex-wrap">
        
          <div className="bg-white flex flex-col w-80 p-8 gap-4 rounded-md shadow-md">
            <h3 className="font-custom">Order Summary</h3>
          
            <div className="flex flex-row justify-between">
              <p>Original price</p>
              <p>$180</p>
            </div>
          <hr className="text-color4"/>

          <div className="flex flex-row justify-between">
              <h3 className="font-custom">Total price</h3>
              <p className="font-custom text-green-600">$180</p>
            </div>

            <button className="p-2 bg-green-700 text-white font-bold rounded-md shadow-md">Proceed to Checkout</button>
            <Link className="underline font-bold flex justify-center text-green-700">
            <p>Continue to shopping</p>
            
            </Link> 
          
            </div>


      
          <div className="bg-white flex flex-col w-80 p-8 gap-4  rounded-md shadow-md">
                      <h3 className="text-dash font-bold "> Do you have a voucher or a gift card?</h3>
                    
                      <input type="text" placeholder="Enter your coupon code here...." className="bg-color4 p-2 text-dash rounded-md shadow-md outline-none font-bold"/>

                      <button className="p-2 bg-green-700 text-white font-bold rounded-md shadow-md">Apply Code</button>
                      <div className="flex flex-col gap-1">
                      <p className="text-dash font-bold text-gray-600">Your available coupon:</p>
                        <p className="text-dash font-bold text-gray-600">XHGJFHE235443 - 10% off </p>
                      </div>
                        
                      </div>

                </div>
    </div>
  )
}

export default OrderSummary
