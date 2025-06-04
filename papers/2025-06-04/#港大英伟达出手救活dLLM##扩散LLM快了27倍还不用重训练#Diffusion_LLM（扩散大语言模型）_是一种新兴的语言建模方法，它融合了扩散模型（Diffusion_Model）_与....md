# #港大英伟达出手救活dLLM##扩散LLM快了27倍还不用重训练#Diffusion LLM（扩散大语言模型） 是一种新兴的语言建模方法，它融合了扩散模型（Diffusion Model） 与...

**URL**: https://weibo.com/6105753431/PuX4budo8

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E6%B8%AF%E5%A4%A7%E8%8B%B1%E4%BC%9F%E8%BE%BE%E5%87%BA%E6%89%8B%E6%95%91%E6%B4%BBdLLM%23&amp;extparam=%23%E6%B8%AF%E5%A4%A7%E8%8B%B1%E4%BC%9F%E8%BE%BE%E5%87%BA%E6%89%8B%E6%95%91%E6%B4%BBdLLM%23" data-hide=""><span class="surl-text">#港大英伟达出手救活dLLM#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E6%89%A9%E6%95%A3LLM%E5%BF%AB%E4%BA%8627%E5%80%8D%E8%BF%98%E4%B8%8D%E7%94%A8%E9%87%8D%E8%AE%AD%E7%BB%83%23&amp;extparam=%23%E6%89%A9%E6%95%A3LLM%E5%BF%AB%E4%BA%8627%E5%80%8D%E8%BF%98%E4%B8%8D%E7%94%A8%E9%87%8D%E8%AE%AD%E7%BB%83%23" data-hide=""><span class="surl-text">#扩散LLM快了27倍还不用重训练#</span></a><br><br>Diffusion LLM（扩散大语言模型） 是一种新兴的语言建模方法，它融合了扩散模型（Diffusion Model） 与传统的大语言模型（LLM） 架构。<br><br>不同于自回归模型逐词生成，扩散模型可以同时生成整体结构，更利于保持句子连贯性，适合诗歌、复杂文案。<br><br>然而，扩散模型通常需要更多训练步骤，经历几十到上百步去噪才能生成文本，相比Transformer那种“一步出一个词”的方式慢很多。<br><br>但最近港大、英伟达和MIT联合提出了Fast-dLLM，实测让这类模型加速多达27.6倍，关键是：完全不需要重新训练。【图1】【图2】<br><br>Fast-dLLM的核心优化可以拆成两块：<br><br>- KV Cache机制重做了：针对双向注意力的Diffusion LLM，Fast-dLLM设计了块级的Key-Value缓存结构，可以缓存多个token的激活值，大幅减少重复计算。而进阶版DualCache，甚至连还没解码的后缀token也能缓存，提高复用率。【图3】<br><br>- 并行解码不再盲猜：Diffusion LLM原本多token解码精度不稳，是因为强行假设token之间独立。Fast-dLLM提出“基于置信度”的并行解码，只在高置信的token上并行生成，其他保留MASK，下一步再解。实测可以从原本64步压缩到37步完成生成，既快又稳。【图4】<br><br>令人惊喜的是，这两套机制可以组合使用：<br><br>- 单独用KV Cache，提升2～3.6倍吞吐量<br><br>- 单独用并行解码，额外提升1.5～2倍<br><br>- 合并使用，GSM8K上加速11倍，MBPP上加速9.2倍<br><br>除了快，该方法的准确率也有保障。比如GSM8K任务中，准确率下降不足2%，仍保持在77%以上。【图5】<br><br>这项方案已经在多个开源Diffusion LLM上验证过，包括LLaDA和Dream模型。适配性强，性能涨幅也更明显出现在“长序列+多shot”的任务上，比如1024 token+8-shot推理。【图6】<br><br>感兴趣的小伙伴可以点击：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fnvlabs.github.io%2FFast-dLLM%2F" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><img style="" src="https://tvax3.sinaimg.cn/large/006Fd7o3gy1i2348yg9ebj31n61ai7qu.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax2.sinaimg.cn/large/006Fd7o3gy1i2348zidy0j316a0rmqes.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax4.sinaimg.cn/large/006Fd7o3gy1i23490ue7zj30zk0cj0z5.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax1.sinaimg.cn/large/006Fd7o3gy1i2349msgy1g33qq18gu12.gif" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax1.sinaimg.cn/large/006Fd7o3gy1i2349csjcuj315g0ii469.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax2.sinaimg.cn/large/006Fd7o3gy1i2349eet44j30zk0jqn6j.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

香港大学、英伟达和MIT联合提出Fast-dLLM技术，通过两项关键创新显著提升扩散大语言模型（Diffusion LLM）效率：1）改进KV缓存机制，采用块级缓存和DualCache技术减少重复计算；2）引入基于置信度的并行解码策略，仅在高可信度token上并行生成。实验显示该方法无需重新训练即可实现最高27.6倍加速，在GSM8K等任务中保持77%以上准确率（仅下降2%），特别适用于长序列多shot场景。该技术已成功适配LLaDA等开源模型，显著提升文本生成效率。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-04T04:05:31Z
- **目录日期**: 2025-06-04
