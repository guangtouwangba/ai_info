# #哈工大点云分析新SOTA##KAN入局3D任务参数量暴降#点云分析圈迎来“新脑子”！哈尔滨工业大学（深圳）和宾夕法尼亚大学联合推出基于Kolmogorov-Arnold Networks...

**URL**: https://weibo.com/6105753431/PsyLicbvR

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E5%93%88%E5%B7%A5%E5%A4%A7%E7%82%B9%E4%BA%91%E5%88%86%E6%9E%90%E6%96%B0SOTA%23&amp;extparam=%23%E5%93%88%E5%B7%A5%E5%A4%A7%E7%82%B9%E4%BA%91%E5%88%86%E6%9E%90%E6%96%B0SOTA%23" data-hide=""><span class="surl-text">#哈工大点云分析新SOTA#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23KAN%E5%85%A5%E5%B1%803D%E4%BB%BB%E5%8A%A1%E5%8F%82%E6%95%B0%E9%87%8F%E6%9A%B4%E9%99%8D%23&amp;extparam=%23KAN%E5%85%A5%E5%B1%803D%E4%BB%BB%E5%8A%A1%E5%8F%82%E6%95%B0%E9%87%8F%E6%9A%B4%E9%99%8D%23" data-hide=""><span class="surl-text">#KAN入局3D任务参数量暴降#</span></a><br><br>点云分析圈迎来“新脑子”！<br><br>哈尔滨工业大学（深圳）和宾夕法尼亚大学联合推出基于Kolmogorov-Arnold Networks（KANs）的新模型PointKAN，在3D感知任务上大放异彩。传统MLPs虽高效，但在捕捉复杂几何时捉襟见肘，而KANs用可学习函数替代固定激活函数，解锁了更灵活的几何特征学习能力。<br><br>PointKAN核心结构包括：<br><br>- Geometric Affine Module：组合Group-Norm与S-Pool，既做归一化又尽量保留信息；<br>- Local Feature Processing：用KAN Block并行处理局部特征，加上深度卷积提升通道信息提取；<br>- Global Feature Processing：靠Residual Point Block叠加提取全局特征。<br><br>升级版PointKAN-elite还引入Efficient-KANs架构，用有理函数替代原本难并行的B样条，大幅减小参数量与计算开销。同时，通过组内参数共享，进一步压缩模型规模。<br><br>实验结果来看，在多个数据集上，PointKAN与PointKAN-elite在分类、分割、小样本任务中全面超越基于MLP的方案<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fmp.weixin.qq.com%2Fs%2F1_sgobR3RgYMzD9QhDacSA" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">参数量暴降，精度反升！哈工大宾大联手打造点云分析新SOTA</span></a> <a href="https://weibo.com/ttarticle/p/show?id=2309405168068124344392" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_article_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">参数量暴降，精度反升！哈工大宾大联手打造点云分析新SOTA</span></a><img style="" src="https://tvax1.sinaimg.cn/large/006Fd7o3gy1i1kxcm654kj30rs0fmtav.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

哈尔滨工业大学（深圳）与宾夕法尼亚大学联合提出基于KANs的新型点云分析模型PointKAN，通过可学习函数替代传统MLP的固定激活函数，显著提升几何特征学习能力。其核心模块包括几何仿射、局部特征并行处理（KAN Block+深度卷积）和全局残差特征提取。升级版PointKAN-elite采用有理函数优化计算效率，结合参数共享实现参数量暴降。实验显示，该模型在分类、分割及小样本任务中全面超越MLP方案，实现精度与效率双突破。相关成果已通过论文及技术文章发布。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-19T11:03:05Z
- **目录日期**: 2025-05-19
