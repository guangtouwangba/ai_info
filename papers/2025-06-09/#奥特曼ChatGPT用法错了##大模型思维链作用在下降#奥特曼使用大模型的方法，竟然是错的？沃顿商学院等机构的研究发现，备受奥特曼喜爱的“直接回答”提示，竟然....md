# #奥特曼ChatGPT用法错了##大模型思维链作用在下降#奥特曼使用大模型的方法，竟然是错的？沃顿商学院等机构的研究发现，备受奥特曼喜爱的“直接回答”提示，竟然...

**URL**: https://weibo.com/6105753431/PvJEUD8Mi

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E5%A5%A5%E7%89%B9%E6%9B%BCChatGPT%E7%94%A8%E6%B3%95%E9%94%99%E4%BA%86%23&amp;extparam=%23%E5%A5%A5%E7%89%B9%E6%9B%BCChatGPT%E7%94%A8%E6%B3%95%E9%94%99%E4%BA%86%23" data-hide=""><span class="surl-text">#奥特曼ChatGPT用法错了#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E5%A4%A7%E6%A8%A1%E5%9E%8B%E6%80%9D%E7%BB%B4%E9%93%BE%E4%BD%9C%E7%94%A8%E5%9C%A8%E4%B8%8B%E9%99%8D%23&amp;extparam=%23%E5%A4%A7%E6%A8%A1%E5%9E%8B%E6%80%9D%E7%BB%B4%E9%93%BE%E4%BD%9C%E7%94%A8%E5%9C%A8%E4%B8%8B%E9%99%8D%23" data-hide=""><span class="surl-text">#大模型思维链作用在下降#</span></a><br><br>奥特曼使用大模型的方法，竟然是错的？<br><br>沃顿商学院等机构的研究发现，备受奥特曼喜爱的“直接回答”提示，竟然会显著降低模型准确率。【图1】<br><br>就比如，“直接回答”和“思维链提示”（CoT）这种，效果其实并不如预期，甚至还有副作用。<br><br>具体来说：<br><br>1. “直接回答”准确率下降<br><br>    - 奥特曼主推的“just answer”策略，研究发现会显著降低模型准确率；<br>    - 特别是在高标准下，表现比“默认模式”差不少。<br>        <br>2. 思维链提示CoT，也没那么神<br><br>    - 对于推理模型（如o3-mini、o4-mini），加入“Think step by step”提示提升非常有限，甚至可能拖慢回答速度；<br>    - o3-mini时间成本+80%，准确率只涨4.1%；<br>    - Gemini 2.5 Flash更是提示越多，效果越差。<br>        <br>3. 非推理模型情况更复杂<br><br>    - CoT对非推理模型有一定提升，但提高的是“平均准确率”，不是稳定性；<br>    - 在“100%准确率”这种严格标准下，部分模型表现反而更差；<br>    - 而且计算成本也在飙升。<br>        <br>4. 默认模式反而是最稳妥的选择<br><br>    - 越来越多模型内置了推理机制和提示模板；<br>    - 这意味着：很多“CoT”、“直接回答”的显式提示，反而可能干扰模型判断；<br>    - 默认模式（无额外提示）+高质量数据，或许是当前最理智的用法。<br>        <br>看来对于直接使用模型应用的用户来说，默认设置就已经是一种很好的使用方式了。<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fmp.weixin.qq.com%2Fs%2FWnSqo7xPxzShz3OGlIxzmA" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">奥特曼ChatGPT用法错了！最新研究：要求“直接回答”降低准确率，思维链提示作用也在下降</span></a><br><br>研究报告地址：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fpapers.ssrn.com%2Fsol3%2Fpapers.cfm%3Fabstract_id%3D5285532" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><img style="" src="https://tvax1.sinaimg.cn/large/006Fd7o3gy1i291pu2bo6j30u00rltfx.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax4.sinaimg.cn/large/006Fd7o3gy1i291uajg7hj319a1aw4gi.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

沃顿商学院最新研究发现，OpenAI CEO奥特曼推崇的"直接回答"提示方式会显著降低大模型准确率，特别是在高标准评估下表现不如默认模式。同时，思维链提示(CoT)对推理模型效果有限，甚至可能增加80%时间成本仅提升4.1%准确率；对非推理模型虽能提高平均准确率，但在严格标准下表现更差且计算成本飙升。研究表明，现代大模型已内置优化机制，使用默认模式配合高质量数据可能是当前最稳妥的使用方式，额外提示反而可能干扰模型性能。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-09T09:02:39Z
- **目录日期**: 2025-06-09
