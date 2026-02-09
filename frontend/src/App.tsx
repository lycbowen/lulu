import { Carousel } from "antd"
import './App.css'
import { GetConfig } from '../wailsjs/go/main/App'
import { useEffect, useState } from "react"
import React from "react"

const images = Object.values(
    import.meta.glob('./assets/images/*.{jpg,jpeg,png}', {
        eager: true,
        import: 'default',
    })
) as string[]

function App() {
    const [fitMode, setFitMode] = useState<React.CSSProperties['objectFit']>('contain');
    const [playSpeed, setPlaySpeed] = useState(5000);

    useEffect(() => {
        GetConfig().then((res) => {
            setFitMode(res.fit_mode as React.CSSProperties['objectFit']);
            setPlaySpeed(res.play_speed);
        })
    }, []);

    return (
        <Carousel className="container" fade speed={1000} autoplay autoplaySpeed={playSpeed} draggable>
            {images.map((img, i) => (
                <div className="slide" key={i}>
                    <img className="pic" src={img} style={{ objectFit: fitMode }} />
                </div>
            ))}
        </Carousel>
    )
}

export default App
