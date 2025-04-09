# #英伟达开源模型超越DeepSeek##英伟达新模型参数量比DeepSeek少一半#英伟达也搞了个开源模型，参数量比DeepSeek R1少一半，性能却实现了超越。这款模型基于Llama...

**URL**: https://weibo.com/6105753431/PmryDi25V

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E8%8B%B1%E4%BC%9F%E8%BE%BE%E5%BC%80%E6%BA%90%E6%A8%A1%E5%9E%8B%E8%B6%85%E8%B6%8ADeepSeek%23&amp;extparam=%23%E8%8B%B1%E4%BC%9F%E8%BE%BE%E5%BC%80%E6%BA%90%E6%A8%A1%E5%9E%8B%E8%B6%85%E8%B6%8ADeepSeek%23" data-hide=""><span class="surl-text">#英伟达开源模型超越DeepSeek#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E8%8B%B1%E4%BC%9F%E8%BE%BE%E6%96%B0%E6%A8%A1%E5%9E%8B%E5%8F%82%E6%95%B0%E9%87%8F%E6%AF%94DeepSeek%E5%B0%91%E4%B8%80%E5%8D%8A%23&amp;extparam=%23%E8%8B%B1%E4%BC%9F%E8%BE%BE%E6%96%B0%E6%A8%A1%E5%9E%8B%E5%8F%82%E6%95%B0%E9%87%8F%E6%AF%94DeepSeek%E5%B0%91%E4%B8%80%E5%8D%8A%23" data-hide=""><span class="surl-text">#英伟达新模型参数量比DeepSeek少一半#</span></a><br><br>英伟达也搞了个开源模型，参数量比DeepSeek R1少一半，性能却实现了超越。<br><br>这款模型基于Llama-3.1训练，名为Llama-3.1-Nemotron-Ultra-253B-v1，以下是该模型的亮点：<br><br>参数规模更小：Llama-3.1-Nemotron-Ultra的参数量为253B，相比DeepSeek R1的671B，小了一半还多。在多项基准测试中，该模型表现出色：在GPQA问答中达到76.01%（对比R1的71.5%）、指令执行任务中取得89.45%（对比R1的83.3%），在LiveCodeBench编程任务中也以66.31%超过R1的65.9%。<br><br>推理架构优化：该模型通过神经架构搜索(NAS, Neural Architecture Search)引入了跳跃注意力层、融合前馈神经网络（FFN）等结构，显著降低了内存占用和计算成本。这些优化使模型可以在搭载8张H100 GPU的服务器上高效运行，降低企业部署门槛。<br><br>支持推理模式：Nemotron Ultra采用多阶段的后训练流程，包括跨领域监督微调（覆盖数学、代码、对话、工具使用等场景）、通过群体相对策略优化(GRPO, Group Relative Policy Optimization)进行强化学习，以及知识蒸馏和额外预训练。模型具备“reasoning on”和“reasoning off”两种模式，开发者可通过系统提示灵活切换。<br><br>测试结果显著：在“推理模式”开启后，模型的数学推理能力提升明显，MATH500得分从80.40%跃升至97.00%，AIME25从16.67%提升至72.50%；在代码任务中，LiveCodeBench分数从29.03%一举升至66.31%。<br><br>多语言多场景：Nemotron Ultra具备多语言能力，支持英语、德语、法语、意大利语、葡萄牙语、印地语、西班牙语、泰语等语言。其适用范围广泛，涵盖对话机器人、AI Agent、检索增强生成(RAG, Retrieval-Augmented Generation)、代码生成等常见大型语言模型应用场景。<br><br>完全开源：英伟达此次不仅开放了模型代码，还公开了权重和后训练数据。不过在商用时，用户需遵守Nvidia Open Model License及Llama 3.1 Community License。英伟达也强调了“负责任AI开发”的重要性，提醒开发者在部署前对模型的对齐、安全性和潜在偏见进行充分评估。<br><br>如果想深入体验，小伙伴们可以直接去Hugging Face下载模型代码和权重：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fhuggingface.co%2Fnvidia%2FLlama-3_1-Nemotron-Ultra-253B-v1%23evaluation-results" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><img style="" src="https://tvax3.sinaimg.cn/large/006Fd7o3gy1i0affpb9lnj325x103nma.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax1.sinaimg.cn/large/006Fd7o3gy1i0affq8ab9j30zk0jg7f5.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

英伟达开源了基于Llama-3.1训练的253B参数模型Llama-3.1-Nemotron-Ultra-253B-v1，参数量仅为DeepSeek R1（671B）的38%，但在GPQA问答（76.01% vs 71.5%）、指令执行（89.45% vs 83.3%）等基准测试中表现更优。该模型通过神经架构搜索优化推理结构，支持8张H100 GPU部署，并具备多语言能力和"推理模式"切换功能（开启后MATH500得分从80.4%提升至97%）。完全开源模型权重及训练数据，适用于对话、代码生成等场景，商用需遵守双许可证协议。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-09T07:02:55Z
- **目录日期**: 2025-04-09
