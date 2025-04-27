# #无损压缩大模型##大模型压缩后体积降到70%# 无损压缩大模型，体积降到70%的那种！你没有听错，在本地部署这一块，动不动几百G的大模型，经常让人头大。而Rice大...

**URL**: https://weibo.com/6105753431/PpdnJ6JZV

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E6%97%A0%E6%8D%9F%E5%8E%8B%E7%BC%A9%E5%A4%A7%E6%A8%A1%E5%9E%8B%23&amp;extparam=%23%E6%97%A0%E6%8D%9F%E5%8E%8B%E7%BC%A9%E5%A4%A7%E6%A8%A1%E5%9E%8B%23" data-hide=""><span class="surl-text">#无损压缩大模型#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E5%A4%A7%E6%A8%A1%E5%9E%8B%E5%8E%8B%E7%BC%A9%E5%90%8E%E4%BD%93%E7%A7%AF%E9%99%8D%E5%88%B070%25%23&amp;extparam=%23%E5%A4%A7%E6%A8%A1%E5%9E%8B%E5%8E%8B%E7%BC%A9%E5%90%8E%E4%BD%93%E7%A7%AF%E9%99%8D%E5%88%B070%25%23" data-hide=""><span class="surl-text">#大模型压缩后体积降到70%#</span></a> <br><br>无损压缩大模型，体积降到70%的那种！<br><br>你没有听错，在本地部署这一块，动不动几百G的大模型，经常让人头大。<br><br>而Rice大学和xMAD.ai团队推出的DFloat11无损压缩框架，能在推理时让LLM体积直接缩小到70%，而且输出结果一模一样，没有一丝精度损失。<br><br>以Llama-3.1-405B为例，原本811.7GB的模型经过DFloat11压缩后，体积缩减到551.22GB，单机8张80GB GPU即可轻松运行。<br><br>压缩后，由于大幅减少了CPU与GPU之间的数据搬运开销，推理速度相比传统方案最高提升了38.8倍。<br><br>DFloat11能够实现无损压缩，主要得益于以下几点：<br><br>- 挖掘了BFloat16的压缩潜力：BFloat16虽为避免数值溢出而设计，但训练完成后的模型中，权重的指数部分非常稀疏，平均只使用了2.6个位的信息量，为压缩留下了巨大的空间。<br><br>- 应用了动态长度编码：类似于Huffman编码，DFloat11对常见权重赋予更短的编码长度，稀有权重则用更长编码，从而实现了最优压缩效果。<br><br>- 专门优化了推理效率：不同于传统Zip等静态压缩方法，DFloat11为GPU推理进行了深度优化，采用小型查找表、双阶段调度机制和块级并行解压等技术，在压缩的同时保持了高吞吐量。<br><br>和常规压缩（比如zip那种）最大的区别，是DFloat11不只是压得小，而是推理时也能高效读取解压，真正做到了部署友好。<br><br>目前代码已经开源，感兴趣的小伙伴可以点击——<br>论文：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Farxiv.org%2Fabs%2F2504.11651" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><br>GitHub：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fgithub.com%2FLeanModels%2FDFloat11" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><img style="" src="https://tvax4.sinaimg.cn/large/006Fd7o3gy1i0vhzymr6mj31co12ek93.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax1.sinaimg.cn/large/006Fd7o3gy1i0vi002km9j30rl0t7x12.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

Rice大学和xMAD.ai团队开发了DFloat11无损压缩框架，可将大模型体积压缩至70%且保持精度无损。以Llama-3.1-405B为例，模型从811.7GB缩减到551.22GB，仅需8张80GB GPU即可运行。该技术利用BFloat16的稀疏特性（仅使用2.6位指数信息），结合动态长度编码和GPU优化设计（查找表、并行解压等），使推理速度提升最高38.8倍。相比传统压缩，DFloat11专为高效推理设计，显著减少CPU-GPU数据传输开销。代码已开源。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-27T17:03:00Z
- **目录日期**: 2025-04-27
