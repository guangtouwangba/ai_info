# #达摩院改进DiT拿下SOTA##开源视觉生成算力减半#算力砍半，视觉生成任务依然SOTA！达摩院在ICLR 2025上抛出的DyDiT架构：通过时间步长与空间区域的智能资源分配...

**URL**: https://weibo.com/6105753431/PoRsl4HQF

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E8%BE%BE%E6%91%A9%E9%99%A2%E6%94%B9%E8%BF%9BDiT%E6%8B%BF%E4%B8%8BSOTA%23&amp;extparam=%23%E8%BE%BE%E6%91%A9%E9%99%A2%E6%94%B9%E8%BF%9BDiT%E6%8B%BF%E4%B8%8BSOTA%23" data-hide=""><span class="surl-text">#达摩院改进DiT拿下SOTA#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E5%BC%80%E6%BA%90%E8%A7%86%E8%A7%89%E7%94%9F%E6%88%90%E7%AE%97%E5%8A%9B%E5%87%8F%E5%8D%8A%23&amp;extparam=%23%E5%BC%80%E6%BA%90%E8%A7%86%E8%A7%89%E7%94%9F%E6%88%90%E7%AE%97%E5%8A%9B%E5%87%8F%E5%8D%8A%23" data-hide=""><span class="surl-text">#开源视觉生成算力减半#</span></a><br><br>算力砍半，视觉生成任务依然SOTA！<br><br>达摩院在ICLR 2025上抛出的DyDiT架构：通过时间步长与空间区域的智能资源分配，将DiT模型的推理算力削减51%，生成速度提升1.73倍，而FID指标几乎无损！<br><br>更惊人的是，这一突破仅需3%的微调成本。<br><br>该方法通过引入动态化调整机制，可精准削减视觉生成任务中50%的推理算力，有效缓解传统扩散模型的计算冗余问题，相关工作已开源。【图1】<br><br>算力砍半效果依然SOTA  <br>DiT架构作为当前主流的生成模型框架，有效实现了图像与视频的可控生成，推动生成式AI走向应用爆发。<br><br>然而，DiT架构的多步生成策略存在推理效率低、算力冗余等问题，在执行视觉生成任务容易造成极高的算力消耗，限制其往更广泛的场景落地。<br><br>业内提出高效采样、特征缓存、模型压缩剪枝等方法尝试解决这一问题，但这些方法均针对静态不变模型，又衍生出潜在的冗余浪费问题。<br><br>达摩院（湖畔实验室）、新加坡国立大学、清华大学等联合研究团队在论文《Dynamic Diffusion Transformer》提出了动态架构DyDiT，能够根据时间步长和空间区域自适应调整计算分配，有效缓解视觉生成任务中的算力消耗问题。【图2】<br><br>具体而言，DyDiT能在简单的时间步长使用较窄的模型宽度，减少计算资源；在空间维度上优先处理含有详细信息的主要对象，减少对背景区域的计算资源分配，提升推理效率与减少计算冗余的同时，保持生成质量。<br><br>使用者更可根据自身的资源限制或者部署要求，灵活调整目标的计算量，DyDiT将自动适配模型参数，实现效果与效率的最佳平衡。<br><br>实验结果表明，DyDiT在多个数据集和生成模型下均表现出高稳定性。<br><br>仅用不到3%的微调成本，将DiT-XL的浮点运算次数（FLOPs）减少了51%，生成速度提高了1.73倍，在ImageNet测得的FID得分与原模型几乎相当（2.27vs2.07）。<br><br>据透露，DyDiT相关训练与推理代码已开源，并计划适配到更多的文生图、文生视频模型上，目前基于知名文生图模型FLUX调试的Dy-FLUX也在开源项目上架。<br><br>据悉，达摩院今年共有13篇论文被ICLR 2025录用，涵盖了视频生成、自然语言处理、医疗AI、基因智能等领域，其中3篇被选为Spotlight。<br><br>论文链接：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Farxiv.org%2Fabs%2F2410.03456" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><br>技术解读：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fmp.weixin.qq.com%2Fs%2FyqYg272vIztflZ6NfX5zJw" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><br>开源链接：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fgithub.com%2Falibaba-damo-academy%2FDyDiT" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><img style="" src="https://tvax3.sinaimg.cn/large/006Fd7o3gy1i0ssw3xo1tj30uk09awgp.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax2.sinaimg.cn/large/006Fd7o3gy1i0ssw6x7brj30jk0n0q9z.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

阿里巴巴达摩院联合多所高校在ICLR 2025提出动态架构DyDiT，通过智能分配计算资源（时间步长动态调整模型宽度、空间维度优先处理主体区域），将DiT模型的推理算力减少51%，生成速度提升1.73倍，同时保持FID指标稳定（2.27 vs 原模型2.07）。该方法仅需3%微调成本，有效解决传统扩散模型算力冗余问题，相关代码已开源并适配文生图模型FLUX。实验显示DyDiT在多个数据集均具高稳定性，为视觉生成任务提供了高效解决方案。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-25T07:02:26Z
- **目录日期**: 2025-04-25
