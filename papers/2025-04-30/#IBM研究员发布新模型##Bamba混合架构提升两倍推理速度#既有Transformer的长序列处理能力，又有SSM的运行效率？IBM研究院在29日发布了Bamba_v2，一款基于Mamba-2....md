# #IBM研究员发布新模型##Bamba混合架构提升两倍推理速度#既有Transformer的长序列处理能力，又有SSM的运行效率？IBM研究院在29日发布了Bamba v2，一款基于Mamba-2...

**URL**: https://weibo.com/6105753431/PpEaLldpy

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23IBM%E7%A0%94%E7%A9%B6%E5%91%98%E5%8F%91%E5%B8%83%E6%96%B0%E6%A8%A1%E5%9E%8B%23&amp;extparam=%23IBM%E7%A0%94%E7%A9%B6%E5%91%98%E5%8F%91%E5%B8%83%E6%96%B0%E6%A8%A1%E5%9E%8B%23" data-hide=""><span class="surl-text">#IBM研究员发布新模型#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23Bamba%E6%B7%B7%E5%90%88%E6%9E%B6%E6%9E%84%E6%8F%90%E5%8D%87%E4%B8%A4%E5%80%8D%E6%8E%A8%E7%90%86%E9%80%9F%E5%BA%A6%23&amp;extparam=%23Bamba%E6%B7%B7%E5%90%88%E6%9E%B6%E6%9E%84%E6%8F%90%E5%8D%87%E4%B8%A4%E5%80%8D%E6%8E%A8%E7%90%86%E9%80%9F%E5%BA%A6%23" data-hide=""><span class="surl-text">#Bamba混合架构提升两倍推理速度#</span></a><br><br>既有Transformer的长序列处理能力，又有SSM的运行效率？<br><br>IBM研究院在29日发布了Bamba v2，一款基于Mamba-2架构的纯解码器语言模型，旨在处理广泛的文本生成任务，其核心特性即将被整合至IBM Granite 4.0当中。<br><br>通过显著降低Transformer键值缓存的内存需求，90亿参数的Bamba模型在保持同等精度下，运行速度可达同类规模Transformer的两倍以上。<br><br>Bamba的最初诞生源于Transformer架构固有的“平方级瓶颈”，生成文本的累积成本呈平方级增长。<br><br>这不仅造成了模型问答的延迟，同时也造成大量冗余计算。早在2022年ChatGPT普及Transformer时，研究者们就已开始寻找替代架构。<br><br>状态空间模型（SSMs）及SSM层与Transformer混合架构，已成为两大潜在解决方案。<br><br>2023年，门控SSM变体Mamba2被提出，推动了一系列混合架构的出现。英伟达去年证实，这些新型混合架构不仅能超越单一架构性能，还能大幅提升推理速度。<br><br>最新推出的Bamba v2相较v1，新增1万亿token训练数据，性能显著提升。<br><br>- 基准测试<br><br>在L1和L2基准测试中，Bamba v2的表现超越了训练数据量近5倍的Llama 3.1 8B模型。<br><br>配合最新vLLM优化，其推理速度达到同规模Transformer模型的2-2.5倍。<br><br>  - HF OpenLLM v1 基准测试（含 OpenbookQA、BoolQ 和 PIQA 评测）【图2】<br>  - HF OpenLLM v2 基准测试【图3】<br><br>- 训练过程<br>鉴于GPU资源有限（仅192块A100显卡），团队采用了两种方案：一是为现有模型注入新数据，二是尝试模型融合技术。【图4】<br><br>训练流程可分为三步：<br>1. 以2万亿token的基础检查点为起点，融入Olmo Mix数据集进行扩展训练，采用恒定学习率2e-5，将训练规模从2万亿token提升至2.5万亿token<br>2. 混合使用Nemotron-CC和Hugging Face数据，继续训练5000亿token使总量达到3万亿。进行恒定学习率和余弦衰减学习率对比实验，验证两种学习率策略的互补效果<br>3. 对两个模型进行1000亿token的高质量数据精炼，并使用MergeKit进行融合，通过加权平均生成最终Bamba 9B v2模型<br><br>团队透露，他们接下来要攻克的难题，是优化vLLM来运行SSM模型。<br><br>技术博客：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fhuggingface.co%2Fblog%2Fibm-ai-platform%2Fbamba-9b-v2" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><br>开源仓库：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fgithub.com%2Ffoundation-model-stack%2Fbamba" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><img style="" src="https://tvax1.sinaimg.cn/large/006Fd7o3gy1i0ys9hybawj31hc0u00w1.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax3.sinaimg.cn/large/006Fd7o3gy1i0ys9kmu3rj313q0pe48f.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax4.sinaimg.cn/large/006Fd7o3gy1i0ys9n2nrzj315k114k44.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax1.sinaimg.cn/large/006Fd7o3gy1i0ys9phk1aj324m0wmgx4.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

IBM研究院发布新型语言模型Bamba v2，基于Mamba-2混合架构，结合Transformer长序列处理能力和状态空间模型(SSM)的高效性。该90亿参数模型通过优化键值缓存内存需求，在保持精度同时实现推理速度达同类Transformer的2-2.5倍。训练采用三阶段策略：基础数据扩展、混合数据增强及模型融合，最终使用3万亿token数据。基准测试显示其性能超越数据量5倍的Llama 3.1 8B模型。团队下一步将优化vLLM以更好支持SSM架构。模型已开源，将整合至IBM Granite 4.0系统。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-30T09:04:05Z
- **目录日期**: 2025-04-30
