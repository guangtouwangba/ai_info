# #一招选出最能推理的数据##小模型教大模型挑数据#和人工标记数据说拜拜，利用预训练语言模型中的注意力机制就能选择可激发推理能力的训练数据！字节Seed团队最新...

**URL**: https://weibo.com/6105753431/PrWX4eIv3

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E4%B8%80%E6%8B%9B%E9%80%89%E5%87%BA%E6%9C%80%E8%83%BD%E6%8E%A8%E7%90%86%E7%9A%84%E6%95%B0%E6%8D%AE%23&amp;extparam=%23%E4%B8%80%E6%8B%9B%E9%80%89%E5%87%BA%E6%9C%80%E8%83%BD%E6%8E%A8%E7%90%86%E7%9A%84%E6%95%B0%E6%8D%AE%23" data-hide=""><span class="surl-text">#一招选出最能推理的数据#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E5%B0%8F%E6%A8%A1%E5%9E%8B%E6%95%99%E5%A4%A7%E6%A8%A1%E5%9E%8B%E6%8C%91%E6%95%B0%E6%8D%AE%23&amp;extparam=%23%E5%B0%8F%E6%A8%A1%E5%9E%8B%E6%95%99%E5%A4%A7%E6%A8%A1%E5%9E%8B%E6%8C%91%E6%95%B0%E6%8D%AE%23" data-hide=""><span class="surl-text">#小模型教大模型挑数据#</span></a><br><br>和人工标记数据说拜拜，利用预训练语言模型中的注意力机制就能选择可激发推理能力的训练数据！<br><br>字节Seed团队最新宣布了一个重要成果——AttentionInfluence。<br><br>无需训练，无需标签，只需用1.3B模型给7B模型选择数据，就能提升模型推理能力，甚至也能提升代码生成能力。<br><br>以往，筛选数据的方法通常依赖于监督分类器，需要人工或大语言模型进行标注，难免引入领域特定偏见。<br><br>字节Seed团队注意到：<br><br>预训练模型中的检索头与检索和上下文推理紧密相关。<br><br>检索头在训练早期就会出现，逐渐增强，并最终在训练的中后期阶段牢固建立，对模型性能起到至关重要的作用。<br><br>1.3B参数稠密模型中检索头的演化过程，be like：【图1】<br><br>但如果直接关闭它们会怎样？<br><br>他们用小型预训练语言模型通过简单的注意力头屏蔽操作，充当强大的模型的数据选择器。<br><br>具体操作是，识别重要检索头，屏蔽这些头以创建性能下降的“弱”模型，计算“弱”模型与原始“强”模型之间的损失差异，根据损失增加幅度对数据进行排名，形成影响分数。<br><br>没想到，实验后他们得到了一个惊人结果。<br><br>将AttentionInfluence方法应用于1.3B参数预训练语言模型，对SmolLM语料库进行数据选择，筛选出73.1B tokens与完整的SmolLM语料库组合，使用WSD方法预训练7B模型。<br><br>在知识密集型和推理密集型基准测试中模型性能均有提升，具体来说：<br><br>MMLU+1.4个百分点、MMLU-Pro+2.7个百分点、AGIEval-en+1.8个百分点、GSM8K+2.7个百分点、HumanEval+3.5个百分点。【图2】<br><br>这项研究发布后引来不少网友关注，谷歌DeepMind研究科学家都转发为其点赞：【图3】<br><br>有网友看后表示：<br><br>多么简单而巧妙的数据选择思路！【图4】<br><br>关于这项研究的更多细节，我们接着往下看：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fmp.weixin.qq.com%2Fs%2FFlP_m6WuWrvxrF4fvgyR9A" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">字节最新大模型秘籍：只挑能有推理潜力的数据训练！1.3B模型无需标签自动挑选</span></a><img style="" src="https://tvax1.sinaimg.cn/large/006Fd7o3gy1i1ga9tjf9fj30pg060wln.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax2.sinaimg.cn/large/006Fd7o3gy1i1ga9vkwaej30q60t6150.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax2.sinaimg.cn/large/006Fd7o3gy1i1ga9xhvvmj30q00f0jwh.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax2.sinaimg.cn/large/006Fd7o3gy1i1ga9zsl65j30qc0im7ff.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

字节Seed团队提出AttentionInfluence方法，利用小模型注意力机制自动筛选能提升大模型推理能力的数据。该方法通过识别并屏蔽1.3B模型中的关键检索头，计算损失差异来评估数据价值，无需人工标注。实验显示，用该方法筛选的73.1B tokens数据训练7B模型后，在MMLU、GSM8K等基准测试中性能显著提升（最高+3.5%）。这种无监督、高效的数据选择方案因创新性获得学界关注，被赞为"简单而巧妙"的突破。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-15T12:04:29Z
- **目录日期**: 2025-05-15
