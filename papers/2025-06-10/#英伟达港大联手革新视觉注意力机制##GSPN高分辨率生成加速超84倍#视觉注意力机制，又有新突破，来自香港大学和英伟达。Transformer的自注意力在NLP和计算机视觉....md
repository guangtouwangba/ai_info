# #英伟达港大联手革新视觉注意力机制##GSPN高分辨率生成加速超84倍#视觉注意力机制，又有新突破，来自香港大学和英伟达。Transformer的自注意力在NLP和计算机视觉...

**URL**: https://weibo.com/6105753431/PvU9CArlG

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E8%8B%B1%E4%BC%9F%E8%BE%BE%E6%B8%AF%E5%A4%A7%E8%81%94%E6%89%8B%E9%9D%A9%E6%96%B0%E8%A7%86%E8%A7%89%E6%B3%A8%E6%84%8F%E5%8A%9B%E6%9C%BA%E5%88%B6%23&amp;extparam=%23%E8%8B%B1%E4%BC%9F%E8%BE%BE%E6%B8%AF%E5%A4%A7%E8%81%94%E6%89%8B%E9%9D%A9%E6%96%B0%E8%A7%86%E8%A7%89%E6%B3%A8%E6%84%8F%E5%8A%9B%E6%9C%BA%E5%88%B6%23" data-hide=""><span class="surl-text">#英伟达港大联手革新视觉注意力机制#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23GSPN%E9%AB%98%E5%88%86%E8%BE%A8%E7%8E%87%E7%94%9F%E6%88%90%E5%8A%A0%E9%80%9F%E8%B6%8584%E5%80%8D%23&amp;extparam=%23GSPN%E9%AB%98%E5%88%86%E8%BE%A8%E7%8E%87%E7%94%9F%E6%88%90%E5%8A%A0%E9%80%9F%E8%B6%8584%E5%80%8D%23" data-hide=""><span class="surl-text">#GSPN高分辨率生成加速超84倍#</span></a><br><br>视觉注意力机制，又有新突破，来自香港大学和英伟达。<br><br>Transformer的自注意力在NLP和计算机视觉领域表现出色——它能捕捉远距离依赖，构建深度上下文。然而，面对高分辨率图像时，传统自注意力有两个大难题：<br><br>- 计算量巨大：O(N²) 的复杂度让处理长上下文变得非常耗时。<br>- 破坏空间结构：将二维图像拉平成一维序列，会丢失像素之间的空间关系。<br><br>虽然线性注意力和Mamba等方法能把复杂度降到O(N)，但它们还是把图像当作一维序列处理，无法真正利用二维空间信息。<br><br>为此，香港大学与英伟达联合推出了广义空间传播网络（GSPN）。<br><br>GSPN采用二维线性传播，结合“稳定性–上下文条件”，将计算量从 O(N²) 或 O(N) 再降到√N级别，并完整保留图像的空间连贯性。这样，不仅大幅提升了效率，还在多个视觉任务上刷新了性能纪录。 <a href="https://weibo.com/ttarticle/p/show?id=2309405176040292221075" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_article_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">英伟达港大联手革新视觉注意力机制！GSPN高分辨率生成加速超84倍</span></a><img style="" src="https://tvax4.sinaimg.cn/large/006Fd7o3gy1i2a9njk3tpj30lj0c4wft.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

香港大学与英伟达联合提出广义空间传播网络（GSPN），革新视觉注意力机制。传统Transformer处理高分辨率图像时存在计算复杂度高（O(N²)）和破坏空间结构的问题。GSPN通过二维线性传播和稳定性-上下文条件，将计算复杂度降至√N级别，同时保留图像空间连贯性。该方法显著提升效率，在高分辨率生成任务中加速超84倍，并在多个视觉任务中刷新性能纪录。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-10T12:03:18Z
- **目录日期**: 2025-06-10
