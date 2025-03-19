# #LLM也能用来做OCR##用Gemma3本地进行OCR#说起OCR（光学字符识别），大多数人第一反应可能是调用云端API。但现在，有人带来了新玩法——利用LLM做OCR，实现完全...

**URL**: https://weibo.com/6105753431/Pj8cs8EGa

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23LLM%E4%B9%9F%E8%83%BD%E7%94%A8%E6%9D%A5%E5%81%9AOCR%23&amp;extparam=%23LLM%E4%B9%9F%E8%83%BD%E7%94%A8%E6%9D%A5%E5%81%9AOCR%23" data-hide=""><span class="surl-text">#LLM也能用来做OCR#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E7%94%A8Gemma3%E6%9C%AC%E5%9C%B0%E8%BF%9B%E8%A1%8COCR%23&amp;extparam=%23%E7%94%A8Gemma3%E6%9C%AC%E5%9C%B0%E8%BF%9B%E8%A1%8COCR%23" data-hide=""><span class="surl-text">#用Gemma3本地进行OCR#</span></a><br><br>说起OCR（光学字符识别），大多数人第一反应可能是调用云端API。但现在，有人带来了新玩法——利用LLM做OCR，实现完全本地化运行。<br><br>看这个【视频】，用户上传一张票据，并点击提取，不一会就在右侧显示了详细的数据，连排版都帮你做了。<br><br>传统OCR主要关注字符识别，而LLM驱动的OCR则更进一步，不仅能提取文字，还能理解图像中的信息结构，比如表格、段落，甚至复杂排版。<br> <br>该方案基于谷歌的多模态模型Gemma-3，结合Streamlit框架，即可实现无需联网条件下的本地OCR处理，大幅提升隐私安全性。<br><br>基于Gemma-3的视觉能力，用户只需依托Ollama环境，并搭配Python，就能直接在本地使用OCR功能。<br><br>相比传统OCR依赖规则或简单机器学习模型，对复杂场景（如手写体、低清扫描件）识别效果一般。<br><br>而Gemma-3 OCR具备大模型的推理能力，能够结合上下文调整识别结果，使OCR不仅能“看见字”，还能“理解内容”，适应各种复杂情况。<br><br>Gemma-3目前提供多个版本，参数规模从10亿到270亿，开发者可根据硬件配置和性能需求选择合适的模型。即使是最小的10亿参数版本，也能支持多模态输入。<br><br>感兴趣的小伙伴可以点击：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fgithub.com%2Fpatchy631%2Fai-engineering-hub%2Ftree%2Fmain%2Fgemma3-ocr" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a> <a href="https://video.weibo.com/show?fid=1034:5145597376659472" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_video_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">量子位的微博视频</span></a><br clear="both"><div style="clear: both"></div><video controls="controls" poster="https://tvax4.sinaimg.cn/orj480/006Fd7o3ly1hzl8tr0mxyj30wq0k074y.jpg" style="width: 100%"><source src="https://f.video.weibocdn.com/o0/sPookLt5lx08mM0dsGc0010412005szo0E010.mp4?label=mp4_720p&amp;template=1178x720.25.0&amp;ori=0&amp;ps=1CwnkDw1GXwCQx&amp;Expires=1742350141&amp;ssig=zyU4%2F2qWUD&amp;KID=unistore,video"><source src="https://f.video.weibocdn.com/o0/4NQ4a2Pmlx08mM0dw19S010412002T7P0E010.mp4?label=mp4_hd&amp;template=784x480.25.0&amp;ori=0&amp;ps=1CwnkDw1GXwCQx&amp;Expires=1742350141&amp;ssig=HgUqCesVHx&amp;KID=unistore,video"><source src="https://f.video.weibocdn.com/o0/6pBNH8qBlx08mM0dnHvG010412001XG70E010.mp4?label=mp4_ld&amp;template=588x360.25.0&amp;ori=0&amp;ps=1CwnkDw1GXwCQx&amp;Expires=1742350141&amp;ssig=3OflFVnxup&amp;KID=unistore,video"><p>视频无法显示，请前往<a href="https://video.weibo.com/show?fid=1034%3A5145597376659472" target="_blank" rel="noopener noreferrer">微博视频</a>观看。</p></video>

## AI 摘要

传统OCR技术主要依赖云端API进行字符识别，而新方法利用大型语言模型（LLM）如Gemma-3实现本地化OCR处理。Gemma-3不仅能提取文字，还能理解图像中的信息结构，如表格和复杂排版，适应各种复杂场景。该方案结合Streamlit框架和Ollama环境，支持多模态输入，提升隐私安全性。Gemma-3提供多个版本，参数规模从10亿到270亿，开发者可根据需求选择合适模型。这种方法相比传统OCR在处理手写体、低清扫描件等复杂场景时表现更优。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-03-19T09:11:29+08:00
