# #AI生成SVG矢量图##AI不只会画图还会写SVG代码#对从事教学图、产品图和数据可视化的人来说，可无限缩放、不失真的SVG矢量图几乎是刚需。而AI模型StarVector，正...

**URL**: https://weibo.com/6105753431/Pk2JM4IUF

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23AI%E7%94%9F%E6%88%90SVG%E7%9F%A2%E9%87%8F%E5%9B%BE%23&amp;extparam=%23AI%E7%94%9F%E6%88%90SVG%E7%9F%A2%E9%87%8F%E5%9B%BE%23" data-hide=""><span class="surl-text">#AI生成SVG矢量图#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23AI%E4%B8%8D%E5%8F%AA%E4%BC%9A%E7%94%BB%E5%9B%BE%E8%BF%98%E4%BC%9A%E5%86%99SVG%E4%BB%A3%E7%A0%81%23&amp;extparam=%23AI%E4%B8%8D%E5%8F%AA%E4%BC%9A%E7%94%BB%E5%9B%BE%E8%BF%98%E4%BC%9A%E5%86%99SVG%E4%BB%A3%E7%A0%81%23" data-hide=""><span class="surl-text">#AI不只会画图还会写SVG代码#</span></a><br><br>对从事教学图、产品图和数据可视化的人来说，可无限缩放、不失真的SVG矢量图几乎是刚需。而AI模型StarVector，正是一款能画出高质量SVG图形的AI工具。<br><br>准确地说，不是画SVG，而是直接写SVG代码，而且写得那叫一个优雅，生产环境可以直接用的那种。<br><br>使用方法也很简单：给它一张图，或者一句文字描述，StarVector就能直接吐出一份标准的SVG代码。<br><br>这段代码采用了如&lt;circle&gt;、&lt;polygon&gt;、&lt;rect&gt;、&lt;text&gt;等SVG原生语义元素，结构简洁，层级分明，视觉与语义兼顾。<br><br>更厉害的是，它支持从文本直接生成SVG。<br><br>比如你输入一句：“一个红色圆圈包着一个写着Start的矩形”，它就能还原出一个 &lt;circle&gt; 包 &lt;rect&gt; 和 &lt;text&gt; 的图形结构。<br><br>StarVector的背后，是一个名为SVG-Stack的专属数据集，包含210万组图像、SVG代码以及自动生成的文字描述。<br><br>这相当于给模型“看”完了一整座互联网的“设计图纸”，几乎任何图形风格它都能快速上手。<br><br>在评价指标上，传统像素级指标如MSE和SSIM在矢量图领域并不适用。<br><br>于是，StarVector自创了DinoScore，用DINOv2模型提取图像语义特征，再计算相似度，而非简单比对像素差异。<br><br>实验表明，DinoScore与人类主观评分高度一致。<br><br>StarVector提供两个版本，满足不同使用场景：<br><br>• StarVector-1B：轻量小模型，适合快速原型；<br>• StarVector-8B：重装旗舰版，搭配StarCoder2和更强图像编码器，能处理更复杂的图，效果也更精准。<br>  <br>性能上，StarVector生成一张图只需 74 秒，对比传统方法动辄20 多分钟，体验简直是降维打击。<br><br>别看它是代码驱动的、没有图像反馈、受限于16k token上下文，这些限制在实力面前都不值一提。<br><br>未来的 SVG 图形，很可能真的不是“画”出来的，而是AI 写出来的。<br><br>感兴趣的小伙伴可以点击：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fgithub.com%2Fjoanrod%2Fstar-vector" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><br>论文：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Farxiv.org%2Fabs%2F2312.11556" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><img style="" src="https://tvax2.sinaimg.cn/large/006Fd7o3gy1hzs6fb5x1ij30zk09o7dc.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax3.sinaimg.cn/large/006Fd7o3gy1hzs6fcxqkjj30zk0gxdth.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

StarVector是一款能直接生成高质量SVG矢量图代码的AI工具。它通过文本描述或输入图像，自动输出结构简洁、可直接用于生产环境的SVG代码（如<circle>、<rect>等原生元素）。其核心技术基于包含210万组数据的SVG-Stack数据集，并创新采用DinoScore语义相似度评估指标。提供1B（轻量版）和8B（精准版）两个版本，生成速度仅74秒，远超传统方法。该工具特别适合教学图、产品图和数据可视化领域，展现了AI通过代码生成矢量图形的强大能力。论文及代码已开源。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-03-25T00:03:00Z
- **目录日期**: 2025-03-25
