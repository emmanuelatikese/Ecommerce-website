import { Swiper, SwiperSlide } from 'swiper/react';
import 'swiper/css';
import 'swiper/css/pagination';
import {FreeMode, Pagination, Scrollbar} from "swiper/modules"
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
        scrollbar={true}
        freeMode={true}
        modules={[Pagination, Scrollbar, FreeMode]}
        className=' w-4/5 h-96 relative top-28'
    >
    <SwiperSlide className=' w-96 h-96 rounded-lg overflow-hidden'>
    <motion.img className='opacity-90 cursor-pointer' whileHover={{scale:1.2, borderRadius:"8px", opacity:0.8, transition:{ ease: easeOut, duration: 0.5}}} 
    exit={{duration:0.5}} src={Trouser} />
    
    <div className='p-4 w-full  bg-opacity-50 bg-black absolute bottom-0 text-white font-bold '>
    <h1 className='text-2xl text-marigold-200'>Pair of trousers</h1>
    <button className='text-white'>View</button>
    </div>
    </SwiperSlide>


    </Swiper>
  )
}

export default slider