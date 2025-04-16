# #MiniMax押注线性注意力##MiniMax让长文本训练速度基本不变#在Transformer几乎一统天下的今天，MiniMax-01选择了另一条路：放弃“主流”Transformer，押注更小众...

**URL**: https://weibo.com/6105753431/PnwYG64wM

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23MiniMax%E6%8A%BC%E6%B3%A8%E7%BA%BF%E6%80%A7%E6%B3%A8%E6%84%8F%E5%8A%9B%23&amp;extparam=%23MiniMax%E6%8A%BC%E6%B3%A8%E7%BA%BF%E6%80%A7%E6%B3%A8%E6%84%8F%E5%8A%9B%23" data-hide=""><span class="surl-text">#MiniMax押注线性注意力#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23MiniMax%E8%AE%A9%E9%95%BF%E6%96%87%E6%9C%AC%E8%AE%AD%E7%BB%83%E9%80%9F%E5%BA%A6%E5%9F%BA%E6%9C%AC%E4%B8%8D%E5%8F%98%23&amp;extparam=%23MiniMax%E8%AE%A9%E9%95%BF%E6%96%87%E6%9C%AC%E8%AE%AD%E7%BB%83%E9%80%9F%E5%BA%A6%E5%9F%BA%E6%9C%AC%E4%B8%8D%E5%8F%98%23" data-hide=""><span class="surl-text">#MiniMax让长文本训练速度基本不变#</span></a><br><br>在Transformer几乎一统天下的今天，MiniMax-01选择了另一条路：放弃“主流”Transformer，押注更小众但计算更高效的线性注意力（linear attention），并将其规模推进到惊人的4560亿参数，搅动开源圈。<br><br>线性注意力是什么？一句话解释，它是一种能将原本计算量为 O(n²) 的attention结构压缩为 O(n) 的优化方法。<br><br>但它很早期时效果并不好，也少有人关注。直到MiniMax团队将其持续打磨，并在2023年内推出了多个关键技术，包括：<br><br>- 用cos函数替代softmax的Cosformer；<br><br>- 分析性能瓶颈的The Devil in Linear Transformer；<br><br>- 更快的Lightning Attention，通过分块算法提升速度、降低延迟；<br><br>- 与Transformer结合的Hybrid架构，既保证性能也保留一定retrieval能力（即上下文记忆）；<br><br>这些研究逐步把线性注意力从“理论好看、实际不行”的尴尬地带，推到了足以工业部署的成熟阶段。<br><br>架构负责人钟怡然说，当大家还在担心线性注意力放大之后会不会失效时，MiniMax已经用400多B规模的模型验证了它“能跑、能快、能长记性”。<br><br>他提到，这种架构的本质优势是随着序列变长，成本优势会越来越大：在1M长度输入下，softmax attention的延迟是lightning attention的2700倍。而这也让lightning架构在长文本生成、长链推理等方向上具备独特优势。<br><br>不过，线性注意力也不是完美。retrieval能力弱是当前难解的瓶颈，这也是目前行业仍偏好hybrid架构的主要原因。未来可能会通过更极致地稀疏化softmax attention，进一步降低成本但保留核心能力。<br><br>钟怡然还透露，下一步他们可能探索的方向是统一理解与生成的大模型架构，即多模态原生模型。他也认为，面向AGI终局，O(n)复杂度的架构才更符合人类智能的认知模式。<br><br>开源地址：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fgithub.com%2FMiniMax-AI%2FMiniMax-01" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><br>模型下载：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fhuggingface.co%2FMiniMaxAI%2FMiniMax-Text-01" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><br>更多细节：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fmp.weixin.qq.com%2Fs%2FNigAnui9fXbfresW8KIX-Q" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">MiniMax押注线性注意力，让百万级长文本只用1/2700算力｜对话MiniMax-01架构负责人钟怡然</span></a><img style="" src="https://tvax4.sinaimg.cn/large/006Fd7o3gy1i0ip2onycaj30my0gqaf5.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax1.sinaimg.cn/large/006Fd7o3gy1i0ip2r0h85j30n20skdop.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax2.sinaimg.cn/large/006Fd7o3gy1i0ip2uscsnj30u00rv7bo.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

MiniMax团队突破性地采用线性注意力(linear attention)架构，将计算复杂度从O(n²)降至O(n)，并成功应用于4560亿参数的大模型。通过关键技术如Cosformer、Lightning Attention等优化，解决了早期性能不足的问题，在长文本处理上展现出显著优势（1M长度时延迟仅为传统方法的1/2700）。虽然仍存在检索能力较弱的局限，但混合架构已实现工业级应用。团队认为O(n)复杂度更符合AGI发展方向，未来将探索多模态原生模型。该成果标志着线性注意力从理论到实践的重要突破。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-16T08:04:11Z
- **目录日期**: 2025-04-16
