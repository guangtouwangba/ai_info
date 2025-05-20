# #一图看懂扩散与流匹配##生成模型过程可视化#两种主流生成模型Flow Matching和Diffusion有啥区别？从结果上看，两种方法都能生成目标图案，比如一个“笑脸”，但...

**URL**: https://weibo.com/6105753431/PsHRf6OnB

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E4%B8%80%E5%9B%BE%E7%9C%8B%E6%87%82%E6%89%A9%E6%95%A3%E4%B8%8E%E6%B5%81%E5%8C%B9%E9%85%8D%23&amp;extparam=%23%E4%B8%80%E5%9B%BE%E7%9C%8B%E6%87%82%E6%89%A9%E6%95%A3%E4%B8%8E%E6%B5%81%E5%8C%B9%E9%85%8D%23" data-hide=""><span class="surl-text">#一图看懂扩散与流匹配#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E7%94%9F%E6%88%90%E6%A8%A1%E5%9E%8B%E8%BF%87%E7%A8%8B%E5%8F%AF%E8%A7%86%E5%8C%96%23&amp;extparam=%23%E7%94%9F%E6%88%90%E6%A8%A1%E5%9E%8B%E8%BF%87%E7%A8%8B%E5%8F%AF%E8%A7%86%E5%8C%96%23" data-hide=""><span class="surl-text">#生成模型过程可视化#</span></a><br><br>两种主流生成模型Flow Matching和Diffusion有啥区别？<br><br>从结果上看，两种方法都能生成目标图案，比如一个“笑脸”，但过程却有显著差异，在【视频1】中：<br><br>- 图像左侧是初始状态，像一个模糊的蓝色团块。<br><br>- 图像右侧是目标分布，也就是我们想生成的图案，比如有“眼睛”和“嘴巴”的笑脸。<br><br>红色粒子代表从起点出发的数据点，它们如何移动，体现了模型生成的“路径”。<br><br>1、Flow Matching（流匹配）：<br><br>- 轨迹线平滑、直接，看起来像被“规划好路线”的导航。<br><br>- 这是因为Flow Matching训练的是一个连续的向量场，告诉粒子“每一步往哪里走最合适”。<br><br>- 这里用的是Euler采样器，结果就是非常光滑的迁移。<br><br>2、Diffusion（扩散模型）：<br><br>- 路径呈现出明显的“抖动”，中间经历了不少波折。<br><br>- 实际上用的是 DDPM（Denoising Diffusion Probabilistic Model），它本质上是在高斯噪声里一步步“去噪”，过程充满随机性。<br><br>- 虽然最终也能收敛到目标分布，但生成路径就显得更“混沌”。<br><br>此外，展示这个过程的可视化工具，支持选择不同采样器，还可以控制时间轴。【视频2】<br><br>感兴趣的小伙伴可以自己尝试一下：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fgithub.com%2Fhelblazer811%2FDiffusion-Explorer" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><img style="" src="https://tvax3.sinaimg.cn/large/006Fd7o3ly1i1m1ot93s6j30vo0k00tx.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax3.sinaimg.cn/large/006Fd7o3ly1i1m1ouy13mj30z80k0gmf.jpg" referrerpolicy="no-referrer"><br><br><br clear="both"><div style="clear: both"></div><video controls="controls" poster="https://tvax4.sinaimg.cn/orj480/006Fd7o3ly1i1m1otdht2j30vo0k00tx.jpg" style="width: 100%"><source src="https://f.video.weibocdn.com/o0/9HgiVostlx08oorx9BSo010412006EdR0E010.mp4?label=mp4_720p&amp;template=1140x720.25.0&amp;ori=0&amp;ps=1CwnkDw1GXwCQx&amp;Expires=1747739389&amp;ssig=hsLv0DE6bl&amp;KID=unistore,video"><source src="https://f.video.weibocdn.com/o0/ve3OZAKtlx08oorx3Z8c010412003Uh00E010.mp4?label=mp4_hd&amp;template=760x480.25.0&amp;ori=0&amp;ps=1CwnkDw1GXwCQx&amp;Expires=1747739389&amp;ssig=G6bygBLBjL&amp;KID=unistore,video"><source src="https://f.video.weibocdn.com/o0/GQBcRkD8lx08oorx89sk010412002INo0E010.mp4?label=mp4_ld&amp;template=568x360.25.0&amp;ori=0&amp;ps=1CwnkDw1GXwCQx&amp;Expires=1747739389&amp;ssig=6aoN%2BinKPW&amp;KID=unistore,video"><p>视频无法显示，请前往<a href="https://video.weibo.com/show?fid=1034%3A5168417615839258" target="_blank" rel="noopener noreferrer">微博视频</a>观看。</p></video>

## AI 摘要

这篇微博对比了两种主流生成模型Flow Matching和Diffusion的生成过程差异。Flow Matching通过连续的向量场规划路径，粒子轨迹平滑直接；而Diffusion模型（如DDPM）通过逐步去噪生成，路径呈现随机抖动。两者最终都能生成目标图案（如笑脸），但Flow Matching过程更有序，Diffusion则更随机。文末提供了可视化工具链接，支持不同采样器和时间轴控制，方便用户观察生成过程动态。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-20T10:10:47Z
- **目录日期**: 2025-05-20
