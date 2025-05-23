# #模型无需配对数据也能互通##所有大规模图像模型的潜在表示竟一致#无需任何配对数据，就能实现文本嵌入的模型空间转换？！曾因llya离职OpenAI，在互联网上掀起讨...

**URL**: https://weibo.com/6105753431/PtavhnzVt

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E6%A8%A1%E5%9E%8B%E6%97%A0%E9%9C%80%E9%85%8D%E5%AF%B9%E6%95%B0%E6%8D%AE%E4%B9%9F%E8%83%BD%E4%BA%92%E9%80%9A%23&amp;extparam=%23%E6%A8%A1%E5%9E%8B%E6%97%A0%E9%9C%80%E9%85%8D%E5%AF%B9%E6%95%B0%E6%8D%AE%E4%B9%9F%E8%83%BD%E4%BA%92%E9%80%9A%23" data-hide=""><span class="surl-text">#模型无需配对数据也能互通#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E6%89%80%E6%9C%89%E5%A4%A7%E8%A7%84%E6%A8%A1%E5%9B%BE%E5%83%8F%E6%A8%A1%E5%9E%8B%E7%9A%84%E6%BD%9C%E5%9C%A8%E8%A1%A8%E7%A4%BA%E7%AB%9F%E4%B8%80%E8%87%B4%23&amp;extparam=%23%E6%89%80%E6%9C%89%E5%A4%A7%E8%A7%84%E6%A8%A1%E5%9B%BE%E5%83%8F%E6%A8%A1%E5%9E%8B%E7%9A%84%E6%BD%9C%E5%9C%A8%E8%A1%A8%E7%A4%BA%E7%AB%9F%E4%B8%80%E8%87%B4%23" data-hide=""><span class="surl-text">#所有大规模图像模型的潜在表示竟一致#</span></a><br><br>无需任何配对数据，就能实现文本嵌入的模型空间转换？！<br><br>曾因llya离职OpenAI，在互联网上掀起讨论飓风的柏拉图表示假说提出：所有足够大规模的图像模型都具有相同的潜在表示。【图1】<br><br>那么是否存在针对文本模型的通用潜在结构呢？<br><br>康奈尔大学现在给出了Plus版答案——vec2vec，首个无监督文本嵌入的跨向量空间转换方法。<br><br>利用共享潜在空间，不仅保留嵌入结构和底层输入语义，还能够反推提取嵌入信息。【图2】<br><br>vec2vec在目标嵌入空间中与真实向量的余弦相似度高达0.92，并在超过8000个随机打乱的嵌入上实现完美匹配，揭示了所有编码器在不同架构或训练数据下都拥有几乎相同的表示形式。【图3】<br><br>具体内容，接下来我们一一拆解：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fmp.weixin.qq.com%2Fs%2FSy0aef6VwCz69Wh3LnWmgg" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">无需数据配对，文本嵌入也能互通？康奈尔重磅研究：所有模型都殊途同归</span></a><img style="" src="https://tvax3.sinaimg.cn/large/006Fd7o3ly1i1pjyopazuj30u00zawo8.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax3.sinaimg.cn/large/006Fd7o3ly1i1pjyosuyvj30u00csaj8.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax1.sinaimg.cn/large/006Fd7o3gy1i1pjyvt42sj30sq08sgoi.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

康奈尔大学提出vec2vec方法，首次实现无需配对数据的无监督文本嵌入空间转换。研究发现，不同架构和训练数据的编码器具有几乎相同的潜在表示形式（余弦相似度达0.92），验证了"柏拉图表示假说"在文本领域的适用性。该方法利用共享潜在空间保留语义结构，并能逆向提取嵌入信息，在8000多个随机嵌入上实现完美匹配。这表明大规模模型可能共享通用潜在结构，为跨模型信息互通提供了新思路。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-23T23:04:02Z
- **目录日期**: 2025-05-23
