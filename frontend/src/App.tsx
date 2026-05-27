import { Carousel } from "antd"
import './App.css'
import { GetConfig, GetImages } from '../wailsjs/go/main/App'
import { useEffect, useState } from "react"
import React from "react"

function App() {
    const [fitMode, setFitMode] = useState<React.CSSProperties['objectFit']>('contain');
    const [playSpeed, setPlaySpeed] = useState(5000);
    const [images, setImages] = useState<string[]>([]);
    const [photoPath, setPhotoPath] = useState('photos');
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        Promise.all([GetConfig(), GetImages()])
            .then(([config, imageList]) => {
                setFitMode(config.fit_mode as React.CSSProperties['objectFit']);
                setPlaySpeed(config.play_speed);
                setPhotoPath(config.photo_path || 'photos');
                setImages(imageList);
            })
            .finally(() => setLoading(false));
    }, []);

    if (loading) {
        return <div className="empty-state">正在加载照片...</div>
    }

    if (images.length === 0) {
        return (
            <div className="empty-state">
                <div className="empty-title">没有找到可播放的图片</div>
                <div className="empty-text">请检查 config.json 中的 photo_path：{photoPath}</div>
            </div>
        )
    }

    return (
        <Carousel className="container" fade speed={1000} autoplay autoplaySpeed={playSpeed} draggable>
            {images.map((img, i) => (
                <div className="slide" key={i}>
                    <img className="pic" src={img} style={{ objectFit: fitMode }} alt="" />
                </div>
            ))}
        </Carousel>
    )
}

export default App
