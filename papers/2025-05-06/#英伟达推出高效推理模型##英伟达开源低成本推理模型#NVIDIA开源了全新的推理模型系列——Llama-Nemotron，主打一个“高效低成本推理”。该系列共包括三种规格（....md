# #英伟达推出高效推理模型##英伟达开源低成本推理模型#NVIDIA开源了全新的推理模型系列——Llama-Nemotron，主打一个“高效低成本推理”。该系列共包括三种规格（...

**URL**: https://weibo.com/6105753431/PqxqZAOMw

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E8%8B%B1%E4%BC%9F%E8%BE%BE%E6%8E%A8%E5%87%BA%E9%AB%98%E6%95%88%E6%8E%A8%E7%90%86%E6%A8%A1%E5%9E%8B%23&amp;extparam=%23%E8%8B%B1%E4%BC%9F%E8%BE%BE%E6%8E%A8%E5%87%BA%E9%AB%98%E6%95%88%E6%8E%A8%E7%90%86%E6%A8%A1%E5%9E%8B%23" data-hide=""><span class="surl-text">#英伟达推出高效推理模型#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E8%8B%B1%E4%BC%9F%E8%BE%BE%E5%BC%80%E6%BA%90%E4%BD%8E%E6%88%90%E6%9C%AC%E6%8E%A8%E7%90%86%E6%A8%A1%E5%9E%8B%23&amp;extparam=%23%E8%8B%B1%E4%BC%9F%E8%BE%BE%E5%BC%80%E6%BA%90%E4%BD%8E%E6%88%90%E6%9C%AC%E6%8E%A8%E7%90%86%E6%A8%A1%E5%9E%8B%23" data-hide=""><span class="surl-text">#英伟达开源低成本推理模型#</span></a><br><br>NVIDIA开源了全新的推理模型系列——Llama-Nemotron，主打一个“高效低成本推理”。<br><br>该系列共包括三种规格（8B、49B、253B），不仅支持主动开启或关闭推理模式，还全部支持商用，覆盖从轻量到旗舰的全场景需求。<br><br>三种模型规格如下：<br><br>- LN-Nano（8B）：轻量小模型，适用于中小型场景，部署成本低，响应速度快；<br><br>- LN-Super（49B）：性能与资源消耗平衡，适合对推理能力和成本都有要求的中型项目；<br><br>- LN-Ultra（253B）：旗舰级大模型，可在8张H100上运行，整体推理能力超过DeepSeek-R1，适配复杂任务和企业级需求。<br><br>Llama-Nemotron在技术层面也有诸多亮点：<br><br>- Puzzle架构重构：该模型采用了全新的“Puzzle”架构，而非传统Transformer，通过去除注意力机制、压缩FFN模块，大幅提升计算效率；<br>    <br>- 超大规模数据训练：训练数据覆盖3300万条高质量合成样本，其中数学和代码类数据占比高达3200万，显著增强模型在逻辑与推理领域的泛化能力；<br>    <br>- 五阶段训练流程：包括架构搜索、蒸馏、持续预训练、有监督微调与强化学习，每一阶段均经过精细设计，推动模型性能逐层跃升；<br>    <br>- 多步推理能力优化：精调过程中以DeepSeek-R1为教师模型，引入RLOO、GRPO等强化学习策略，在保持泛化能力的同时显著提升对话质量；<br>    <br>- 基础设施深度优化：训练与推理过程中结合cudagraph和FP8推理技术，使GPU利用率稳定超过90%，大幅降低单位推理成本。<br><br>在多个权威评测基准上，Llama-Nemotron系列在推理与通用理解任务中表现出色，如GPQA、AIME、MATH500等测试成绩，均属领先水平。<br><br>目前，英伟达已将该系列模型连同代码和数据集一同开源，感兴趣的小伙伴可以点击——<br><br>模型：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fhuggingface.co%2Fcollections%2Fnvidia%2Fllama-nemotron-67d92346030a2691293f200b" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><br>论文：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Farxiv.org%2Fabs%2F2505.00949" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><br>Dataset：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fhuggingface.co%2Fdatasets%2Fnvidia%2FLlama-Nemotron-Post-Training-Dataset" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><img style="" src="https://tvax4.sinaimg.cn/large/006Fd7o3gy1i15k8sci5aj31ls0tkk8n.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

英伟达开源了高效低成本推理模型系列Llama-Nemotron，包含8B、49B和253B三种规格，覆盖轻量到旗舰级需求。该系列采用创新的Puzzle架构，去除注意力机制并压缩FFN模块，提升计算效率。训练数据包含3300万高质量样本，强化逻辑推理能力。通过五阶段训练流程和强化学习优化，模型在GPQA等测试中表现优异。结合cudagraph和FP8技术，GPU利用率超90%，显著降低推理成本。所有模型均支持商用，代码和数据集已开源。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-06T05:03:19Z
- **目录日期**: 2025-05-06
