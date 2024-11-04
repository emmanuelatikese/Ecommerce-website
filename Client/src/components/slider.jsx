import { Swiper, SwiperSlide } from 'swiper/react';
import 'swiper/css';
import 'swiper/css/pagination';
import {Pagination} from "swiper/modules"
import Trouser from "../assets/images/trouser.jpg"
import {easeOut, motion} from "framer-motion"

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
    <SwiperSlide className=' w-96 h-96 rounded-lg overflow-hidden'>
    <motion.img whileHover={{scale:1.2, borderRadius:"8px", opacity:0.8, transition:{ ease: easeOut, duration: 0.5}}} 
    exit={{duration:0.5}} src={Trouser} />
    </SwiperSlide>
   
    </Swiper>
  )
}

export default slider