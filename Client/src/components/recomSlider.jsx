import { Swiper, SwiperSlide } from 'swiper/react';
import 'swiper/css';
import 'swiper/css/pagination';
import {Pagination, FreeMode} from "swiper/modules";
import Trouser from "../assets/images/trouser.jpg"
import { Link } from 'react-router-dom';
import ProdIcon from "../icons/product.png"

const AlsoBoughtSlider = () => {
  return (
<Swiper
        pagination={{
        dynamicBullets: true,
        clickable: true,
        }}
        FreeMode={true}
        breakpoints={
            {
                720:{
                    spaceBetween:40,
                    slidesPerView: 3,

                }
            }
        }
        modules={[Pagination, FreeMode]}
        className='flex items-center'
    >

    <SwiperSlide className='flex flex-col gap-2 justify-center items-center'>
     <img  src={Trouser} className='w-96 h-32'/>
    
    </SwiperSlide>







    </Swiper>
  )
}

export default AlsoBoughtSlider