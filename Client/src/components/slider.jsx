import { Swiper, SwiperSlide } from 'swiper/react';
import 'swiper/css';
import 'swiper/css/pagination';
import {Pagination} from "swiper/modules"
import Trouser from "../assets/images/trouser.jpg"

const slider = () => {
  return (
    <Swiper
        pagination={{
        dynamicBullets: true,
        clickable: true
        }}
        breakpoints={
            {
                720:{
                    spaceBetween:40,
                    slidesPerView: 3,

                }
            }
        }
        modules={[Pagination]}
        className=' w-4/5 h-96 relative top-28'
    >
    <SwiperSlide className='w-96 h-72 flex flex-wrap rounded-lg flex-col border-solid border-2'>
    <img src={Trouser}/>
    </SwiperSlide>
   
    </Swiper>
  )
}

export default slider