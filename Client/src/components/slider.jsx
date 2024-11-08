import { Swiper, SwiperSlide } from 'swiper/react';
import 'swiper/css';
import 'swiper/css/pagination';
import {FreeMode, Pagination, Scrollbar} from "swiper/modules"
import Trouser from "../assets/images/trouser.jpg"
import {easeOut, keyframes, motion} from "framer-motion"
import { useState } from 'react'

const Slider = () => {
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
    <SwiperSlide className=' w-96 h-96 rounded-lg overflow-hidden flex flex-col flex-wrap'>
    <motion.img className='opacity-90 cursor-pointer' 
    whileHover={{scale:1.2, borderRadius:"8px", opacity:0.8, transition:{ type: keyframes, ease: easeOut, duration: 0.5,}}} 
    src={Trouser} />
    
    <motion.div
    className='p-4 w-full  bg-opacity-50 bg-black absolute bottom-0 text-white font-bold '>
    <h1 className='text-2xl text-marigold-200'>Pair of trousers</h1>
    <button className='text-white'>View</button>
    </motion.div>
    </SwiperSlide>


    </Swiper>
  )
}

export default Slider