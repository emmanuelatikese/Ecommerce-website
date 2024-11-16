import { Swiper, SwiperSlide } from 'swiper/react';
import 'swiper/css';
import 'swiper/css/pagination';
import {Pagination, FreeMode} from "swiper/modules";
import Trouser from "../assets/images/trouser.jpg"
import { Link } from 'react-router-dom';
import ProdIcon from "../icons/product.png"


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
        modules={[Pagination, FreeMode]}
        className=' w-4/5 flex items-center'
    >
    <SwiperSlide className='w-96 flex p-2 flex-col item-center'>
        <div className='flex flex-col gap-3 justify-center items-center'>
                <img src={Trouser} className='h-96 w-80 rounded-lg shadow-slate-400 shadow-lg'/>
            <div className='w-48 bg-white flex flex-col p-4 rounded-lg items-center shadow-slate-400 shadow-lg gap-4'>
                <div className='font-bold flex flex-row gap-4 items-center'> <img src={ProdIcon} className='w-8 h-8'/> TROUSER</div>
                <Link>
                            <div className='text-white bg-color2 shadow-md text-dash p-2 w-20 rounded-md text-center font-bold'> EXPLORE </div>
                </Link>

            </div>
        </div>
    </SwiperSlide>

    <SwiperSlide className='w-96 flex p-2 flex-col item-center'>
        <div className='flex flex-col gap-3 justify-center items-center'>
                <img src={Trouser} className='h-96 w-80 rounded-lg shadow-slate-400 shadow-lg'/>
            <div className='w-48 bg-white flex flex-col p-4 rounded-lg items-center shadow-slate-400 shadow-lg gap-4'>
                <div className='font-bold flex flex-row gap-4 items-center'> <img src={ProdIcon} className='w-8 h-8'/> TROUSER</div>
                <Link>
                            <div className='text-white bg-color2 shadow-md text-dash p-2 w-20 rounded-md text-center font-bold'> EXPLORE </div>
                </Link>

            </div>
        </div>
    </SwiperSlide>

    <SwiperSlide className='w-96 flex p-2 flex-col item-center'>
        <div className='flex flex-col gap-3 justify-center items-center'>
                <img src={Trouser} className='h-96 w-80 rounded-lg shadow-slate-400 shadow-lg'/>
            <div className='w-48 bg-white flex flex-col p-4 rounded-lg items-center shadow-slate-400 shadow-lg gap-4'>
                <div className='font-bold flex flex-row gap-4 items-center'> <img src={ProdIcon} className='w-8 h-8'/> TROUSER</div>
                <Link>
                            <div className='text-white bg-color2 shadow-md text-dash p-2 w-20 rounded-md text-center font-bold'> EXPLORE </div>
                </Link>

            </div>
        </div>
    </SwiperSlide>

    <SwiperSlide className='w-96 flex p-2 flex-col item-center'>
        <div className='flex flex-col gap-3 justify-center items-center'>
                <img src={Trouser} className='h-96 w-80 rounded-lg shadow-slate-400 shadow-lg'/>
            <div className='w-48 bg-white flex flex-col p-4 rounded-lg items-center shadow-slate-400 shadow-lg gap-4'>
                <div className='font-bold flex flex-row gap-4 items-center'> <img src={ProdIcon} className='w-8 h-8'/> TROUSER</div>
                <Link>
                            <div className='text-white bg-color2 shadow-md text-dash p-2 w-20 rounded-md text-center font-bold'> EXPLORE </div>
                </Link>

            </div>
        </div>
    </SwiperSlide>

    <SwiperSlide className='w-96 flex p-2 flex-col item-center'>
        <div className='flex flex-col gap-3 justify-center items-center'>
                <img src={Trouser} className='h-96 w-80 rounded-lg shadow-slate-400 shadow-lg'/>
            <div className='w-48 bg-white flex flex-col p-4 rounded-lg items-center shadow-slate-400 shadow-lg gap-4'>
                <div className='font-bold flex flex-row gap-4 items-center'> <img src={ProdIcon} className='w-8 h-8'/> TROUSER</div>
                <Link>
                            <div className='text-white bg-color2 shadow-md text-dash p-2 w-20 rounded-md text-center font-bold'> EXPLORE </div>
                </Link>

            </div>
        </div>
    </SwiperSlide>

    </Swiper>
  )
}

export default Slider