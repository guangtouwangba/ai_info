# #陈丹琦在ICLR上主题演讲##如何在学术预算下训练语言模型#计算资源不足、数据访问受限……学术研究者该怎样训练语言模型？一直以来，训练语言模型这一领域几乎完...

**URL**: https://weibo.com/6105753431/Ppdnu4LvZ

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E9%99%88%E4%B8%B9%E7%90%A6%E5%9C%A8ICLR%E4%B8%8A%E4%B8%BB%E9%A2%98%E6%BC%94%E8%AE%B2%23&amp;extparam=%23%E9%99%88%E4%B8%B9%E7%90%A6%E5%9C%A8ICLR%E4%B8%8A%E4%B8%BB%E9%A2%98%E6%BC%94%E8%AE%B2%23" data-hide=""><span class="surl-text">#陈丹琦在ICLR上主题演讲#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E5%A6%82%E4%BD%95%E5%9C%A8%E5%AD%A6%E6%9C%AF%E9%A2%84%E7%AE%97%E4%B8%8B%E8%AE%AD%E7%BB%83%E8%AF%AD%E8%A8%80%E6%A8%A1%E5%9E%8B%23&amp;extparam=%23%E5%A6%82%E4%BD%95%E5%9C%A8%E5%AD%A6%E6%9C%AF%E9%A2%84%E7%AE%97%E4%B8%8B%E8%AE%AD%E7%BB%83%E8%AF%AD%E8%A8%80%E6%A8%A1%E5%9E%8B%23" data-hide=""><span class="surl-text">#如何在学术预算下训练语言模型#</span></a><br><br>计算资源不足、数据访问受限……学术研究者该怎样训练语言模型？<br><br>一直以来，训练语言模型这一领域几乎完全由工业界主导。<br><br>它依靠海量计算资源驱动，并遵循奖励更大模型与数据集的scaling law。<br><br>但对于学术研究者而言，参与其中常令人望而却步。<br><br>就在昨天，陈丹琦在ICLR 2025上进行了名为《学术界训练语言模型：挑战还是使命？》的主题演讲，分享了自己在过去两年间的研究成果。【图1】<br><br>总结而言，她提出了学术研究者能做出重要贡献的三个方向：<br>1. 开发小而精的模型<br>2. 理解与改进训练数据<br>3. 基于开源权重模型发展后训练方法。<br><br>她列举了几项团队的成果进行说明：<br><br>一、剪枝优化LLaMA【图2】<br><br>基于现有LLM进行结构化剪枝。通过复用LLaMA-7B模型的子网络架构和执行持续预训练优化，最终仅使用了相当于从头训练的3%的计算成本，就实现了更好的性能效果。<br><br>二、QuRater评分模型【图3】<br><br>高质量预训练数据的选择对构建强大语言模型至关重要。通过训练QuRater评分模型，将成对判断转化为标量评分，并据此对2600亿token的训练语料进行四维质量标注。<br><br>在实验中，团队根据不同质量评分筛选出300亿token的数据，并基于这些数据训练了13亿参数的语言模型。研究发现，平衡数据质量与多样性至关重要。<br><br>三、ProLong模型【图4】<br><br>ProLong是一个支持512K token上下文长度的80亿参数开源模型。通过持续训练与监督微调技术，提升了长上下文信息的利用效率。<br><br>最终推出的ProLong-8B模型基于Llama-3初始化，使用400亿token训练，在128K上下文长度下达到同规模模型的顶尖水平。尽管长上下文训练数据量仅为Llama-3.1-8B-Instruct的5%，ProLong在多数长上下文任务中表现更优。<br><br>在演讲的最后，她总结道：仅需数百GPU小时，就能开展严肃的语言模型训练研究。甚至偶尔能打造顶尖水平的模型！<br><br>她希望，意识到这一点能够促进学术界更广泛参与语言模型训练，并推动产学研合作的新范式。<img style="" src="https://tvax3.sinaimg.cn/large/006Fd7o3gy1i0vhz99eijj31k0160b29.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax1.sinaimg.cn/large/006Fd7o3gy1i0vhz6k3enj31ts10mk7a.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax4.sinaimg.cn/large/006Fd7o3gy1i0vhz8ach3j31kg0wydyf.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax3.sinaimg.cn/large/006Fd7o3gy1i0vhz9c6elj31lo0yqar5.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax1.sinaimg.cn/large/006Fd7o3gy1i0vhzdevbxj31js0tmn8o.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

陈丹琦在ICLR 2025演讲中提出学术界训练语言模型的三大方向：1）开发高效小模型，如通过剪枝LLaMA-7B仅用3%计算成本实现更好性能；2）优化训练数据，其团队开发的QuRater模型可评估数据质量，用300亿精选token训练出优质13亿参数模型；3）发展后训练方法，如ProLong-8B模型通过微调技术实现512K长上下文处理，性能超越同类。她强调仅需数百GPU小时即可开展有效研究，呼吁学术界更积极参与语言模型开发，推动产学研合作创新。（99字）

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-27T22:03:21Z
- **目录日期**: 2025-04-27
